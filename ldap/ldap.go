package ldap

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"log"
	//"log"
)

//获取Ldap服务器的主机名(或者IP)和端口号
const LdapPort = "636"       //对应代码中的config.Ldap().Port
const LdapHost = "localhost" //对应config.Ldap().Host
//获取Ldap目录结构
const LdapBaseDn = "ou=staff,dc=wang,dc=com"

//获取Ldap数据结构（键值关系）
//在项目中需要获取的ldap数据字段包括：uid(用户id) 、deportment(部门)，displayName(姓名)、mail(邮箱)
const LdapUsernameKey = "uid"          //对应config.Ldap().Attributes.UNameKey
const LdapNameKey = "displayName"      //对应config.Ldap().Attributes.NameKey
const LdapEmailKey = "mail"            //对应config.Ldap().Attributes.EmailKey
const LdapDepartmentKey = "department" //对应config.Ldap().Attributes.DepartmentKey
func ldap_use() {

	/*
		//在项目中进行连接
		//在项目中单独使用一个函数来验证用户名和密码，验证成功则返回一个符合本项目数据库要求的user实体，如中途出现err或者没有拿到数据等都会返回err或者空数据。
		//创建与ldap服务器的链接
		conn, err := ldap.DialTLS("tcp", LdapHost+":"+LdapPort, &tls.Config{InsecureSkipVerify: true})
		if err !=nil{
			panic(err)
		}
		//设置超时时间
		conn.SetTimeout(5 * time.Second)
		defer conn.Close()

		//这里的username为用户点击ldap登录时输入的用户名。
		filter := fmt.Sprintf("(%s=%s)", LdapUsernameKey, LdapNameKey )
		attributes := []string{LdapUsernameKey, LdapDepartmentKey, LdapEmailKey , LdapNameKey}
		sql := ldap.NewSearchRequest(
			//config.Ldap().LoginDn: ou=staff,dc=ebupt,dc=com
			LdapBaseDn,
			//scope:  查询的范围
			ldap.ScopeWholeSubtree,
			//DerefAiases： 在搜索中别名(cn, ou)是否废弃
			ldap.NeverDerefAliases,
			//SizeLimit: 大小设置,一般设置为0
			0,
			//TimeLimit: 时间设置,一般设置为0
			0,
			//TypesOnly:  设置false（返回的值要多一点）
			false,
			//Filter 是过滤条件
			filter,
			//Attributes 需要返回的属性值
			attributes,
			//Controls:  控制
			nil)
	*/
	//连接ldap服务器
	conn, err := ldap.Dial("tcp", "192.168.199.199:389")
	if err != nil {
		panic(err)
	}
	fmt.Println("连接成功")

	//认证
	err = conn.Bind("cn=admin,dc=wjyl,dc=com", "201314")
	if err != nil {
		panic(err)
	}
	fmt.Println("认证成功")

	//查询
	//parmeter1: BaseDN
	//parmeter2: 查询的范围
	//parmeter3:  在搜索中别名(cn, ou)是否废弃
	//SizeLimit: 大小设置,一般设置为0
	//TimeLimit: 时间设置,一般设置为0
	//TypesOnly:  设置false（好像返回的值要多一点）
	//Controls:  是控制我没怎么用过,一般设置nil

	srsql := ldap.NewSearchRequest("ou=user,dc=wjyl,dc=com",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(&(objectClass=inetOrgPerson))",
		[]string{"cn", "sn", "Email"},
		nil)

	cur, er := conn.Search(srsql)

	if er != nil {
		log.Fatalln(er)
	}
	fmt.Println("查询成功")

	if len(cur.Entries) > 0 {

		for _, item := range cur.Entries {

			cn := item.GetAttributeValue("cn")

			fmt.Print(cn + "\t")

			sn := item.GetAttributeValue("sn")

			fmt.Print(sn + "\t")

			objectClass := item.GetAttributeValue("objectClass")
			if objectClass == "" {
				fmt.Println("属性为空")
			} else {
				fmt.Println(objectClass)
			}

		}
	}

	sql := ldap.NewAddRequest("cn=lisi,ou=user,dc=wjyl,dc=com", nil)

	sql.Attribute("uidNumber", []string{"1010"})
	sql.Attribute("gidNumber", []string{"1003"})
	sql.Attribute("userPassword", []string{"123456"})
	sql.Attribute("homeDirectory", []string{"/home/lisi"})
	sql.Attribute("cn", []string{"lisi"})
	sql.Attribute("uid", []string{"lisi"})
	sql.Attribute("objectClass", []string{"shadowAccount", "posixAccount", "account"})
	err = conn.Add(sql)

	if err != nil {
		panic(err)
	}
	fmt.Println("添加新用户成功")

	/*
		client := &ldap.LDAPClient{
		        Base:         "cn=admin,dc=testing,dc=io",
		        Host:         "52.51.245.219",
		        Port:         389,
		        UseSSL:       false,
		        BindDN:       "cn=admin,dc=testing,dc=io",
		        BindPassword: "test123",
		        UserFilter:   "(uid='*api*')",
		        // GroupFilter:  "(memberUid=%s)",
		        Attributes: []string{"givenName", "sn", "mail", "uid"},
		    }
		    defer client.Close()
		    username := "cn=admin,dc=testing,dc=io"
		    password := "test123"
		    ok, user, err := client.Authenticate(username, password)
		    if err != nil {
		        log.Fatalf("Error authenticating user %s: %+v", "*cn=admin,dc=testing,dc=io*", err)
		    }
		    if !ok {
		        log.Fatalf("Authenticating failed for user %s", "*cn=admin,dc=testing,dc=io*")
		    }
		    log.Printf("User: %+v", user)
	*/

}

/*
type ldapService struct{

}
type User struct{
	Id   int
	Password   string
	RoleName   string
	RoleId     int
	LastLogin  int
	HaveLdap   int
	Name    string
	Username  string
	Email     string
	Department  string
}
func (ldapService *ldapService) Login(username string, password string) (*datamodels.User, error) {
	var (
		filter     string
		attributes []string
		conn       *ldap.Conn
		err        error
		cur        *ldap.SearchResult
		user      User
	)

	filter = fmt.Sprintf("(%s=%s)", LdapUsernameKey, LdapNameKey )
	attributes = []string{LdapUsernameKey, LdapDepartmentKey, LdapEmailKey , LdapNameKey}
	if LdapPort == "636" {
		conn, err = ldap.DialTLS("tcp", LdapHost+":"+LdapPort, &tls.Config{InsecureSkipVerify: true})
	} else {
		conn, err = ldap.Dial("tcp", LdapHost+":"+LdapPort)
	}

	if err != nil {
		return nil, err
	}
	conn.SetTimeout(5 * time.Second)
	defer conn.Close()
	sql := ldap.NewSearchRequest(
		LdapBaseDn,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		filter,
		attributes,
		nil)
	if cur, err = conn.Search(sql); err != nil {
		return nil, err
	}
	if len(cur.Entries) == 0 {
		return nil, nil
	}
	entry := cur.Entries[0]
	user =User{
		Id:         utils.GetMillisecond(),
		Password:   config.DefaultPwd,
		RoleName:   config.OrdinaryRole,
		RoleId:     config.OrdinaryRoleId,
		LastLogin:  utils.JsonTime(time.Now()),
		HaveLdap:   config.One,
		Name:       entry.GetAttributeValue(config.Ldap().Attributes.NameKey),
		Username:   entry.GetAttributeValue(config.Ldap().Attributes.UNameKey),
		Email:      entry.GetAttributeValue(config.Ldap().Attributes.EmailKey),
		Department: entry.GetAttributeValue(config.Ldap().Attributes.DepartmentKey),
	}
	err = conn.Bind(entry.DN, password)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	return &user, nil
}
*/

//参考链接
//https://www.cnblogs.com/rongfengliang/p/13659051.html
//https://blog.csdn.net/weixin_35396246/article/details/112630502
//https://github.com/go-ldap/ldap
//https://blog.csdn.net/Mr_rsq/article/details/118937775
//https://baijiahao.baidu.com/s?id=1620805608101105264&wfr=spider&for=pc
