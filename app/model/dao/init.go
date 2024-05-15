package dao

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stone2401/light-gateway/app/tools/db"
)

// 要被初始化的 dao
var syncDao = []any{&Admin{}, &App{}, &ServiceAccessControl{}, &ServiceInfo{},
	&ServiceGrpcRule{}, &ServiceHttpRule{}, &ServiceTcpRule{}, &ServiceLoadBalance{}, &Qps{},
}

// 初始化数据函数
// var initFunc = []func(){initAdmin, initApp, initServiceAccessControl,
// 	initServiceInfo, initServiceGrpcRule, initServiceHttpRule, initServiceLoadBalance,
// 	initServiceTcpRule,
// }

func Init() {
	// 同步表
	go func() {
		// 最大连接数设置
		SyncTable()
		// for _, funcItem := range initFunc {
		// 	funcItem()
		// }
	}()
}

// 同步表结构
func SyncTable() {
	err := db.GetDBDriver().Sync2(syncDao...)
	if err != nil {
		log.Panicln("[i] 数据库表初始化失败: ", err)
	}
}

// 初始化 admin 表
func initAdmin() {
	admin := &Admin{
		Id:       1,
		UserName: "admin",
		Password: "8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918",
		Salt:     "2823d896e9822c0833d41d4904f0c00756d718570fce49b9a379a62c804689d3",
		IsDelete: 0,
	}
	insertTable("admin", admin)
}

// 初始化 app 表
func initApp() {
	apps := []App{
		{Id: 1, AppId: "app_id_a", Name: "租户A", Secret: "449441eb5e72dca9c42a12f3924ea3a2", WhiteIps: "", Qpd: 100000, Qps: 100},
		{Id: 2, AppId: "app_id_b", Name: "租户B", Secret: "8d7b11ec9be0e59a36b52f32366c09cb", WhiteIps: "", Qpd: 0, Qps: 0},
		{Id: 3, AppId: "app_id", Name: "租户名称", Secret: "", WhiteIps: "", Qpd: 0, Qps: 0},
		{Id: 4, AppId: "app_id45", Name: "名称", Secret: "07d980f8a49347523ee1d5c1c41aec02", WhiteIps: "", Qpd: 0, Qps: 0},
	}
	for _, app := range apps {
		insertTable("app", &app)

	}
}

func initServiceAccessControl() {
	serviceAccessControls := []ServiceAccessControl{
		{162, 35, true, "", "", "", 0, 0},
		{165, 34, false, "", "", "", 0, 0},
		{167, 36, false, "", "", "", 0, 0},
		{168, 38, true, "111.11", "22.33", "11.11", 12, 12},
		{169, 41, true, "111.11", "22.33", "11.11", 12, 12},
		{170, 42, true, "111.11", "22.33", "11.11", 12, 12},
		{171, 43, false, "111.11", "22.33", "11.11", 12, 12},
		{172, 44, false, "", "", "", 0, 0},
		{173, 45, false, "", "", "", 0, 0},
		{174, 46, false, "", "", "", 0, 0},
		{175, 47, false, "", "", "", 0, 0},
		{176, 48, false, "", "", "", 0, 0},
		{177, 49, false, "", "", "", 0, 0},
		{178, 50, false, "", "", "", 0, 0},
		{179, 51, false, "", "", "", 0, 0},
		{180, 52, false, "", "", "", 0, 0},
		{181, 53, false, "", "", "", 0, 0},
		{182, 54, true, "127.0.0.3", "127.0.0.2", "", 11, 12},
		{183, 55, true, "127.0.0.2", "127.0.0.1", "", 45, 34},
		{184, 56, false, "192.168.1.0", "", "", 0, 0},
		{185, 57, false, "", "127.0.0.1,127.0.0.2", "", 0, 0},
		{186, 58, true, "", "", "", 0, 0},
		{187, 59, true, "127.0.0.1", "", "", 2, 3},
		{188, 60, true, "", "", "", 0, 0},
		{189, 61, false, "", "", "", 0, 0},
	}
	for _, serviceAccessControl := range serviceAccessControls {
		insertTable("service_access_control", &serviceAccessControl)
	}
}

func initServiceInfo() {
	serviceInfos := []ServiceInfo{
		{Id: 34, LoadType: 0, ServiceName: "websocket_test", ServiceDesc: "websocket_test"},
		{Id: 35, LoadType: 1, ServiceName: "test_grpc", ServiceDesc: "test_grpc"},
		{Id: 36, LoadType: 2, ServiceName: "test_httpe", ServiceDesc: "test_httpe"},
		{Id: 38, LoadType: 0, ServiceName: "service_name", ServiceDesc: "11111"},
		{Id: 41, LoadType: 0, ServiceName: "service_name_tcp", ServiceDesc: "11111"},
		{Id: 42, LoadType: 0, ServiceName: "service_name_tcp2", ServiceDesc: "11111"},
		{Id: 43, LoadType: 1, ServiceName: "service_name_tcp4", ServiceDesc: "service_name_tcp4"},
		{Id: 44, LoadType: 0, ServiceName: "websocket_service", ServiceDesc: "websocket_service"},
		{Id: 45, LoadType: 1, ServiceName: "tcp_service", ServiceDesc: "tcp_desc"},
		{Id: 46, LoadType: 1, ServiceName: "grpc_service", ServiceDesc: "grpc_desc"},
		{Id: 47, LoadType: 0, ServiceName: "testsefsafs", ServiceDesc: "werrqrr"},
		{Id: 48, LoadType: 0, ServiceName: "testsefsafs1", ServiceDesc: "werrqrr"},
		{Id: 49, LoadType: 0, ServiceName: "testsefsafs1222", ServiceDesc: "werrqrr"},
		{Id: 50, LoadType: 2, ServiceName: "grpc_service_name", ServiceDesc: "grpc_service_desc"},
		{Id: 51, LoadType: 2, ServiceName: "gresafsf", ServiceDesc: "wesfsf"},
		{Id: 52, LoadType: 2, ServiceName: "gresafsf11", ServiceDesc: "wesfsf"},
		{Id: 53, LoadType: 2, ServiceName: "tewrqrw111", ServiceDesc: "123313"},
		{Id: 54, LoadType: 2, ServiceName: "test_grpc_service1", ServiceDesc: "test_grpc_service1"},
		{Id: 55, LoadType: 1, ServiceName: "test_tcp_service1", ServiceDesc: "redis服务代理"},
		{Id: 56, LoadType: 0, ServiceName: "test_http_service", ServiceDesc: "测试HTTP代理"},
		{Id: 57, LoadType: 1, ServiceName: "test_tcp_service", ServiceDesc: "测试TCP代理"},
		{Id: 58, LoadType: 2, ServiceName: "test_grpc_service", ServiceDesc: "测试GRPC服务"},
		{Id: 59, LoadType: 0, ServiceName: "test.com:8080", ServiceDesc: "测试域名接入"},
		{Id: 60, LoadType: 0, ServiceName: "test_strip_uri", ServiceDesc: "测试路径接入"},
		{Id: 61, LoadType: 0, ServiceName: "test_https_server", ServiceDesc: "测试https服务"},
	}
	for _, serviceInfo := range serviceInfos {
		insertTable("serviceInfo", &serviceInfo)
	}
}

func initServiceGrpcRule() {
	serviceGrpcRules := []ServiceGrpcRule{
		{171, 53, 8009, ""},
		{172, 54, 8002, "add metadata1 datavalue,edit metadata2 datavalue2"},
		{173, 58, 8012, "add meta_name meta_value"},
	}
	for _, serviceGrpcRule := range serviceGrpcRules {
		insertTable("serviceGrpcRules", &serviceGrpcRule)
	}
}
func initServiceHttpRule() {
	serviceHttpRules := []ServiceHttpRule{
		{165, 35, 8080, 1, "", false, false, false, "", ""},
		{168, 34, 8080, 0, "", false, false, false, "", ""},
		{170, 36, 8080, 0, "", false, false, false, "", ""},
		{171, 38, 8080, 0, "/abc", true, false, true, "^/abc $1", "add head1 value1"},
		{172, 43, 8080, 0, "/usr", true, true, false, "^/afsaasf $1,^/afsaasf $1", ""},
		{173, 44, 8080, 1, "www.test.com", true, true, true, "", ""},
		{174, 47, 8080, 1, "www.test.com", true, true, true, "", ""},
		{175, 48, 8080, 1, "www.test.com", true, true, true, "", ""},
		{176, 49, 8080, 1, "www.test.com", true, true, true, "", ""},
		{177, 56, 8080, 0, "/test_http_service", true, true, true, "^/test_http_service/abb/{.*} /test_http_service/bba/$1", "add header_name header_value"},
		{178, 59, 8080, 1, "test.com", false, true, true, "", "add headername headervalue"},
		{179, 60, 8080, 0, "/test_strip_uri", false, true, false, "^/aaa/{.*} /bbb/$1", ""},
		{180, 61, 8080, 0, "/test_https_server", true, true, false, "", ""},
	}
	for _, serviceHttpRule := range serviceHttpRules {
		insertTable("serviceHttpRule", &serviceHttpRule)
	}
}

func initServiceLoadBalance() {
	serviceLoadBalances := []ServiceLoadBalance{
		{162, 35, 0, 2000, 5000, 2, "127.0.0.1:50051", "100", "", 10000, 0, 0, 0},
		{165, 34, 0, 2000, 5000, 2, "100.90.164.31:8072,100.90.163.51:8072,100.90.163.52:8072,100.90.165.32:8072", "50,50,50,80", "", 20000, 20000, 10000, 100},
		{167, 36, 0, 2000, 5000, 2, "100.90.164.31:8072,100.90.163.51:8072,100.90.163.52:8072,100.90.165.32:8072", "50,50,50,80", "100.90.164.31:8072,100.90.163.51:8072", 10000, 10000, 10000, 100},
		{168, 38, 0, 0, 0, 1, "111:111,22:111", "11,11", "111", 1111, 111, 222, 333},
		{169, 41, 0, 0, 0, 1, "111:111,22:111", "11,11", "111", 0, 0, 0, 0},
		{170, 42, 0, 0, 0, 1, "111:111,22:111", "11,11", "111", 0, 0, 0, 0},
		{171, 43, 0, 2, 5, 1, "111:111,22:111", "11,11", "", 1111, 2222, 333, 444},
		{172, 44, 0, 2, 5, 2, "127.0.0.1:8076", "50", "", 0, 0, 0, 0},
		{173, 45, 0, 2, 5, 2, "127.0.0.1:88", "50", "", 0, 0, 0, 0},
		{174, 46, 0, 2, 5, 2, "127.0.0.1:8002", "50", "", 0, 0, 0, 0},
		{175, 47, 0, 2, 5, 2, "12777:11", "11", "", 0, 0, 0, 0},
		{176, 48, 0, 2, 5, 2, "12777:11", "11", "", 0, 0, 0, 0},
		{177, 49, 0, 2, 5, 2, "12777:11", "11", "", 0, 0, 0, 0},
		{178, 50, 0, 2, 5, 2, "127.0.0.1:8001", "50", "", 0, 0, 0, 0},
		{179, 51, 0, 2, 5, 2, "1212:11", "50", "", 0, 0, 0, 0},
		{180, 52, 0, 2, 5, 2, "1212:11", "50", "", 0, 0, 0, 0},
		{181, 53, 0, 2, 5, 2, "1111:11", "111", "", 0, 0, 0, 0},
		{182, 54, 0, 2, 5, 1, "127.0.0.1:80", "50", "", 0, 0, 0, 0},
		{183, 55, 0, 2, 5, 3, "127.0.0.1:81", "50", "", 0, 0, 0, 0},
		{184, 56, 0, 2, 5, 2, "127.0.0.1:2003,127.0.0.1:2004", "50,50", "", 0, 0, 0, 0},
		{185, 57, 0, 2, 5, 2, "127.0.0.1:6379", "50", "", 0, 0, 0, 0},
		{186, 58, 0, 2, 5, 2, "127.0.0.1:50055", "50", "", 0, 0, 0, 0},
		{187, 59, 0, 2, 5, 2, "127.0.0.1:2003,127.0.0.1:2004", "50,50", "", 0, 0, 0, 0},
		{188, 60, 0, 2, 5, 2, "127.0.0.1:2003,127.0.0.1:2004", "50,50", "", 0, 0, 0, 0},
		{189, 61, 0, 2, 5, 2, "127.0.0.1:3003,127.0.0.1:3004", "50,50", "", 0, 0, 0, 0},
	}
	for _, serviceLoadBalance := range serviceLoadBalances {
		insertTable("serviceLoadBalance", &serviceLoadBalance)
	}
}
func initServiceTcpRule() {
	serviceTcpRules := []ServiceTcpRule{
		{171, 41, 8002},
		{172, 42, 8003},
		{173, 43, 8004},
		{174, 38, 8004},
		{175, 45, 8001},
		{176, 46, 8005},
		{177, 50, 8006},
		{178, 51, 8007},
		{179, 52, 8008},
		{180, 55, 8010},
		{181, 57, 8011},
	}
	for _, serviceTcpRule := range serviceTcpRules {
		insertTable("serviceTcpRule", &serviceTcpRule)
	}
}

// 初始化用到的函数
func insertTable(table string, data interface{}) {
	ok, err := db.GetDBDriver().Get(data)
	if err != nil {
		log.Panicf("[!] 初始化数据表失败 %s 表: %s\n", table, err)
	}
	ok2, _ := db.GetDBDriver().Unscoped().Get(data)
	if !ok && !ok2 {
		i, err2 := db.GetDBDriver().Insert(data)
		if err2 != nil {
			log.Panicf("[!] 初始化数据失败 %s 表: %s\n", table, err2)
		}
		log.Printf("[+] %s 表数据初始化完成，插入 %d 条数据\n", table, i)
	}
}
