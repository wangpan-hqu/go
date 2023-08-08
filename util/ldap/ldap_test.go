package ldap

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"testing"
)

func TestLdap(t *testing.T) {
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

	conn := Connect("192.168.199.199", "389", "admin", "123456")

	fmt.Println(AuthenticateUser(conn, "cn=admin,dc=wjyl,dc=com", "wp", "wjyl123456"))

	sql := ldap.NewAddRequest("cn=lisi,ou=user,dc=wjyl,dc=com", nil)

	sql.Attribute("uidNumber", []string{"1010"})
	sql.Attribute("gidNumber", []string{"1003"})
	sql.Attribute("userPassword", []string{"123456"})
	sql.Attribute("homeDirectory", []string{"/home/lisi"})
	sql.Attribute("cn", []string{"lisi"})
	sql.Attribute("uid", []string{"lisi"})
	sql.Attribute("objectClass", []string{"shadowAccount", "posixAccount", "account"})
	err := conn.Add(sql)
	if err != nil {
		panic(err)
	}
	fmt.Println("添加新用户成功")
}
