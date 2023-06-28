package auth

// Role User
var (
	RoleSuperAdmin        = `super-admin`
	RoleAdminPartner      = `admin-partner`
	RoleCustomerPersonal  = `customer-personal`
	RoleCustomerCorporate = `customer-corporate`
	RoleDriverPersonal    = `driver-personal`
	RoleDriverCorporate   = `driver-corporate`

	RoleSuperAdminView        = `SUPER ADMIN`
	RoleAdminPartnerView      = `ADMIN PARTER`
	RoleDriverPersonalView    = `DRIVER PERSONAL`
	RoleDriverCorporateView   = `DRIVER CORPORATE`
	RoleCustomerPersonalView  = `CUSTOMER PERSONAL`
	RoleCustomerCorporateView = `CUSTOMER CORPORATE`

	RoleSuperAdminCode        = uint64(1)
	RoleAdminPartnerCode      = uint64(2)
	RoleDriverPersonalCode    = uint64(3)
	RoleDriverCorporateCode   = uint64(4)
	RoleCustomerPersonalCode  = uint64(5)
	RoleCustomerCorporateCode = uint64(6)

	UserRolesCode = map[string]*uint64{
		RoleSuperAdmin:        &RoleSuperAdminCode,
		RoleAdminPartner:      &RoleAdminPartnerCode,
		RoleCustomerPersonal:  &RoleDriverPersonalCode,
		RoleCustomerCorporate: &RoleDriverCorporateCode,
		RoleDriverPersonal:    &RoleCustomerPersonalCode,
		RoleDriverCorporate:   &RoleCustomerCorporateCode,
	}

	UserRoles = map[string]string{
		RoleSuperAdmin:        RoleSuperAdmin,
		RoleAdminPartner:      RoleAdminPartner,
		RoleCustomerPersonal:  RoleCustomerPersonal,
		RoleCustomerCorporate: RoleCustomerCorporate,
		RoleDriverPersonal:    RoleDriverPersonal,
		RoleDriverCorporate:   RoleDriverCorporate,
	}

	UserRolesView = map[string]*string{
		RoleSuperAdmin:        &RoleSuperAdminView,
		RoleAdminPartner:      &RoleAdminPartnerView,
		RoleCustomerPersonal:  &RoleCustomerPersonalView,
		RoleCustomerCorporate: &RoleCustomerCorporateView,
		RoleDriverPersonal:    &RoleDriverPersonalView,
		RoleDriverCorporate:   &RoleDriverCorporateView,
	}
)

// Customer Type
var (
	CustomerTypes = map[string]*string{
		RoleCustomerPersonal:  &RoleCustomerPersonal,
		RoleCustomerCorporate: &RoleCustomerCorporate,
	}

	CustomerTypesView = map[string]*string{
		RoleCustomerPersonal:  &RoleCustomerPersonalView,
		RoleCustomerCorporate: &RoleCustomerCorporateView,
	}
)

// Customer Request Status
var (
	RequestStatusWaitConfirmation = "waiting confirmation"
	RequestStatusApproved         = "approved"
	RequestStatusRejected         = "rejected"
)
