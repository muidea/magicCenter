package model

import (

)

type Resource interface {
	RId() int
	RName() string
	RType() int
	
	RRelative() []Resource
}

