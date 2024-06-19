package main

type membershipType string

const (
	TypeStandard membershipType = "standard"
	TypePremium  membershipType = "premium"
)

// don't touch above this line

type Membership struct{
  Type membershipType
  MessageCharLimit int
}

type User struct {
	Name string
  Membership
}

func newUser(name string, membershipType membershipType) User {
  if membershipType == TypePremium{
    return User{
      Name : name,
      Membership : Membership{
        Type : membershipType,
        MessageCharLimit : 1000,
      },
    }
  }else {
    return User{
      Name : name,
      Membership : Membership{
        Type : membershipType,
        MessageCharLimit : 100,
      },
    }
  }
  
  
}

