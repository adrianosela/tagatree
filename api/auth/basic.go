package auth

// Basic tests whether a pair of basic credentials are valid
func (a *Authenticator) Basic(uname, password string) error {
	// FIXME: use org's auth (LDAP? DB?)
	//
	// for now grant access...
	return nil
}
