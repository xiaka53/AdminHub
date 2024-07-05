package mysql

import (
	"fmt"
	"github.com/e421083458/gorm"
	"github.com/xiaka53/AdminHub/dao"
	"golang.org/x/crypto/bcrypt"
	"os"
	"sync"
	"time"
)

type MysqlConf struct {
	Url      string
	Port     int
	User     string
	Pass     string
	Database string
}

func (s *MysqlConf) confTesting() {
	if s.Url == "" {
		s.Url = "127.0.0.1"
	}
	if s.Port == 0 {
		s.Port = 3306
	}
	if s.User == "" {
		s.User = "root"
	}
}

func MysqlInit(sqlConf MysqlConf, name string) {
	(&sqlConf).confTesting()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", sqlConf.User, sqlConf.Pass, sqlConf.Url, sqlConf.Port, sqlConf.Database)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("mysql链接失败")
		os.Exit(0)
	}
	defer db.Close()

	err = db.DB().Ping()
	if err != nil {
		fmt.Println("mysql链接失败")
		os.Exit(0)
	}

	err = db.AutoMigrate(
		dao.Setting{},
		dao.Admin{},
		dao.Menu{},
		dao.Role{},
		dao.UploadSet{},
		dao.UploadFiles{},
		dao.AdminLog{},
	).Error
	if err != nil {
		fmt.Println("mysql链接失败")
		os.Exit(0)
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go sqlConf.writeToml(&wg)
	go sqlInit(db, name, &wg)
	wg.Wait()
}

func sqlInit(db *gorm.DB, name string, w *sync.WaitGroup) {
	defer w.Done()
	baseUserPass := []byte("123456")
	password, err := bcrypt.GenerateFromPassword(baseUserPass, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("管理员密码生成失败")
		os.Exit(0)
	}
	admin := dao.Admin{
		Username:      "admin",
		Password:      string(password),
		LastLoginTime: time.Now(),
		RoleId:        1,
		Atime:         time.Now(),
		Avatar:        "",
	}
	menus := []dao.Menu{
		{1, 0, "基本管理", "", "", 1, "icon-jibenguanli", 1, 0, 1},
		{2, 1, "角色列表", "RoleList", "/admin/role/roleList", 2, "", 0, 0, 1},
		{3, 1, "管理员列表", "AdminList", "/admin/admin/adminList", 2, "", 2, 0, 1},
		{4, 3, "添加管理员", "", "/admin/admin/addAdmin", 3, "", 0, 0, 0},
		{5, 0, "系统设置", "", "", 1, "icon-xitongshezhi", 10, 0, 1},
		{6, 2, "添加角色", "", "/admin/role/addRole", 3, "", 2, 0, 0},
		{7, 2, "编辑角色", "", "/admin/role/editRole", 3, "", 3, 1, 0},
		{8, 2, "删除角色", "", "/admin/role/delRole", 3, "", 4, 0, 0},
		{9, 3, "编辑管理员", "", "/admin/admin/editAdmin", 3, "", 1, 0, 0},
		{10, 3, "删除管理员", "", "/admin/admin/delAdmin", 3, "", 3, 0, 0},
		{11, 1, "操作日志", "OperationLog", "/admin/admin/adminLog", 2, "", 3, 0, 1},
		{12, 5, "菜单管理", "MenuSet", "/admin/menu/menuList", 2, "", 0, 0, 1},
		{13, 12, "新增菜单", "", "/admin/menu/addMenu", 3, "", 0, 0, 0},
		{14, 12, "编辑菜单", "", "/admin/menu/editMenu", 3, "", 1, 0, 0},
		{15, 12, "删除菜单", "", "/admin/menu/delMenu", 3, "", 2, 0, 0},
		{16, 12, "设置日志打印", "", "/admin/menu/setNeedLog", 3, "", 3, 0, 0},
		{17, 5, "上传设置", "UploadSet", "/admin/setting/getUploadConfig", 2, "", 1, 0, 1},
		{18, 5, "参数配置", "BasicInfo", "/admin/setting/settingList", 2, "", 2, 0, 1},
		{19, 18, "新增配置", "", "/admin/setting/addSetting", 3, "", 1, 0, 0},
		{20, 18, "修改配置", "", "/admin/setting/editSetting", 3, "", 3, 0, 0},
		{21, 18, "删除配置", "", "/admin/setting/delSetting", 3, "", 2, 0, 0},
		{22, 17, "删除文件", "", "/admin/setting/delFile", 3, "", 1, 0, 0},
	}
	role := dao.Role{
		RoleId:   1,
		RoleName: "超级管理员",
		Ids:      "[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22]",
		Describe: "最高权限",
		Atime:    time.Now(),
	}
	setting := dao.Setting{
		Title:  dao.SystemName,
		Type:   1,
		Value:  name,
		CanDel: 0,
	}
	uploadSet := dao.UploadSet{
		Qiniu:   "",
		Alioss:  "",
		Txcos:   "",
		Visible: dao.UploadSetVisibleLoc,
	}
	_db := db.Begin()
	if err = _db.Create(&admin).Error; err != nil {
		_db.Rollback()
		fmt.Println("管理员创建失败,err:", err)
		os.Exit(0)
	}
	for _, v := range menus {
		if err = _db.Create(&v).Error; err != nil {
			_db.Rollback()
			fmt.Println("权限列表创建失败,err:", err)
			os.Exit(0)
		}
	}
	if err = _db.Create(&role).Error; err != nil {
		_db.Rollback()
		fmt.Println("权限组创建失败,err:", err)
		os.Exit(0)
	}
	if err = _db.Create(&setting).Error; err != nil {
		_db.Rollback()
		fmt.Println("系统名称设置失败,err:", err)
		os.Exit(0)
	}
	if err = _db.Create(&uploadSet).Error; err != nil {
		_db.Rollback()
		fmt.Println("上传配置设置失败,err:", err)
		os.Exit(0)
	}
	_db.Commit()
	fmt.Println(fmt.Sprintf("系统创建成功：\n登陆用户：%s \n登陆密码：%s", admin.Username, string(baseUserPass)))
}
