package ldap

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"time"
)

func Connect(host, port, username, password string) *ldap.Conn {
	//连接ldap服务器
	conn, err := ldap.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println("ldap Server connection failed", err)
		panic(err)
	}
	/*
		conn, err := ldap.DialTLS("tcp", host+":"+port, &tls.Config{InsecureSkipVerify: true})
		if err !=nil{
			panic(err)
		}
	*/
	conn.SetTimeout(5 * time.Second)
	defer conn.Close()

	err = conn.Bind(username, password)
	if err != nil {
		fmt.Println("Fail to authenticate ldap admin account", err)
		panic(err)
	}
	return conn
}

func AuthenticateUser(conn *ldap.Conn, baseDn, username string, password string) bool {

	//先对用户进行搜索，是否在ldap中
	//person, intetOrgPersopn,
	// filter := fmt.Sprintf("(&(objectClass=inetOrgPerson)(cn=%s))", username)
	filter := fmt.Sprintf("(sAMAccountName=%s)", username)
	//srsql := ldap.NewSearchRequest(localconfig.ExtendConfig.Ldap.BaseDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, filter, []string{"cn", "sn", "mail"}, nil)

	//查询
	//parmeter1: BaseDN
	//parmeter2: 查询的范围
	//parmeter3:  在搜索中别名(cn, ou)是否废弃
	//SizeLimit: 大小设置,一般设置为0
	//TimeLimit: 时间设置,一般设置为0
	//TypesOnly:  设置false（好像返回的值要多一点）
	//Controls:  是控制我没怎么用过,一般设置nil
	srsql := ldap.NewSearchRequest(baseDn, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, filter, []string{"dn", "cn"}, nil)
	cur, err := conn.Search(srsql)

	if err != nil {
		fmt.Println("Fail to search ldap user", err)
		return false
	}
	if len(cur.Entries) < 1 {
		fmt.Println("Cannot locate user information for filter")
		return false

	} else if len(cur.Entries) > 1 {
		fmt.Println("ldap user search found more than one result")
		return false
	}

	//进行二次bind，验证用户password是否正确
	err = conn.Bind(cur.Entries[0].DN, password)
	if err != nil {
		fmt.Println("Username or Password is not Correct", err)
		return false
	}

	return true
}

//参考链接
//https://www.cnblogs.com/rongfengliang/p/13659051.html
//https://blog.csdn.net/weixin_35396246/article/details/112630502
//https://github.com/go-ldap/ldap
//https://blog.csdn.net/Mr_rsq/article/details/118937775
//https://baijiahao.baidu.com/s?id=1620805608101105264&wfr=spider&for=pc
