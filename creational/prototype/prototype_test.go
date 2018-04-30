package prototype

import "testing"

func TestClone(t *testing.T) {
	shirtCache := GetShirtsCloner()
	if shirtCache == nil {
		t.Fatal("received cache was nil")
	}

	item1, err := shirtCache.GetClone(White)
	if err != nil {
		t.Fatal(err)
	}
	if item1 == whitePrototype {
		t.Errorf("item1 cannot be equal to the white prototype")
	}
	shirt1, ok := item1.(*Shirt)
	if !ok {
		t.Fatal("type assertion for shirt1 couldn't be done successfully")
	}
	shirt1.SKU = "abbcc"

	item2, err := shirtCache.GetClone(White)
	if err != nil {
		t.Fatal(err)
	}
	shirt2, ok := item2.(*Shirt)
	if !ok {
		t.Fatal("type assertion for shirt2 couldn't be done succesfully")
	}

	if shirt1.SKU == shirt2.SKU {
		t.Error("shirt 1 cannot be equal to shirt 2")
	}

	t.Logf("LOG: %s", shirt1.GetInfo())
	t.Logf("LOG: %s", shirt2.GetInfo())

	t.Logf("LOG: The mamory positions of the shirts are different %p != %p \n\n", &shirt1, &shirt2)
}
