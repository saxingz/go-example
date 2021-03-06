package my_map

import "testing"

/**
@author saxing 2020/9/28 22:21
*/
func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1 = %d", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2 = %d", len(m2))
	m3 := make(map[int]int, 10)
	t.Logf("len m3 = %d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])

	m1[2] = 0
	t.Log(m1[2])
	t.Log(m1[5])

	if v, ok := m1[8]; ok {
		t.Logf("key value is %d", v)
	} else {
		t.Log("key is not existing...")
	}

	m1[8] = 0

	if v, ok := m1[8]; ok {
		t.Logf("key value is %d", v)
	} else {
		t.Log("key is not existing...")
	}

}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3, 4: 8}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
