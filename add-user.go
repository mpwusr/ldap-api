package ldap_api

import (
	"fmt"
	"github.com/go-ldap/ldap"
	"log"
)

func createDisabledUser() {
	addReq = ldp.NewAddRequest("CN=fooUser,OU=Users,dc=example,dc=com", []ldp.Control{})
	addReq.Attribute("objectClass", []string{"top", "organizationalPerson", "user", "person"})
	addReq.Attribute("name", []string{"fooUser"})
	addReq.Attribute("sAMAccountName", []string{"fooUser"})
	addReq.Attribute("userAccountControl", []string{fmt.Sprintf("%d", 0x0202})
	addReq.Attribute("instanceType", []string{fmt.Sprintf("%d", 0x00000004})
	addReq.Attribute("userPrincipalName", []string{"fooUser@example.com"})
	addReq.Attribute("accountExpires", []string{fmt.Sprintf("%d", 0x00000000})

	addReq.Attributes = attrs

	if err := l.AddRequest(addReq); err != nil {
		log.Fatal("error adding service:", addReq, err)
	}
}

func setUserPassword() {
	// connect code comes here

	// https://github.com/golang/text
	// According to the MS docs the password needs to be enclosed in quotes o_O
	utf16 := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	pwdEncoded, err := utf16.NewEncoder().String(fmt.Sprintf("%q", userPasswd))
	if err != nil {
		log.Fatal(err)
	}

	modReq := ldap.NewModifyRequest("CN=fooUser,OU=Users,dc=example,dc=com", []ldap.Control{})
	modReq.Replace("unicodePwd", []string{pwdEncoded})

	if err := l.ModRequest(modReq); err != nil {
		log.Fatal("error setting user password:", modReq, err)
	}
}

func enableUserAccount() {
	modReq := ldap.NewModifyRequest("CN=fooUser,OU=Users,dc=example,dc=com", []ldap.Control{})
	modReq.Replace("userAccountControl", []string{fmt.Sprintf("%d", 0x0200})

	if err := l.ModRequest(modReq); err != nil {
		log.Fatal("error enabling user account:", modReq, err)
	}
}

func modifyPasswordChoose() {
	passwdModReq := ldap.NewPasswordModifyRequest("", "OldPassword", "NewPassword")
	if _, err = l.PasswordModify(passwdModReq); err != nil {
		log.Fatalf("failed to modify password: %v", err)
	}
}

func modifyPasswordRandom() {
	passwdModReq := ldap.NewPasswordModifyRequest("", "OldPassword", "")
	passwdModResp, err := l.PasswordModify(passwdModReq)
	if err != nil {
		log.Fatalf("failed to change password: %v", err)
	}

	newPasswd := passwdModResp.GeneratedPassword
	log.Printf("New password: %s\n", newPasswd)
}

func deleteRecord() {
	delReq = ldap.NewDelRequest("CN=fooUser,OU=Users,dc=example,dc=com", []ldap.Control{})

	if err := l.Delete(delReq); err != nil {
		log.Fatalf("Error deleting service: %v", err)
	}
}