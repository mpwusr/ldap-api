ldapadd -x -D cn=admin,dc=ldap-domain,dc=com -W -f topdn.ldif
ldapadd -x -D cn=admin,dc=ldap-domain,dc=com -W -f basedb.ldif
ldapadd -x -D cn=admin,dc=ldap-domain,dc=com -W -f ldapusers.ldif
ldapadd -x -D cn=admin,dc=ldap-domain,dc=com -W -f ldapgroups.ldif
