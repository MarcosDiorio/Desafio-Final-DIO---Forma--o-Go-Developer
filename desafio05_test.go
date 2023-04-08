package main

import "testing"

func TestSumCorrect(t *testing.T) {
	teste := soma(88, 1, 741)
	resultado := 830

	if teste != resultado {
		t.Error("Valor esperado :", resultado, "Valor retornado:", teste)
	}
}
func TestSumIncorrect(t *testing.T) {
	teste := soma(88, 1, 741)
	resultado := 555

	if teste != resultado {
		t.Error("Valor esperado :", resultado, "Valor retornado:", teste)
	}
}
func TestMultCorrect(t *testing.T) {
	teste := multiplica(9, 150)
	resultado := 1350
	if teste != resultado {
		t.Error("Valor esperado :", resultado, "Valor retornado:", teste)
	}
}
func TestMultIncorrect(t *testing.T) {
	teste := multiplica(9, 150)
	resultado := 101
	if teste != resultado {
		t.Error("Valor esperado :", resultado, "Valor retornado:", teste)
	}
}
func TestDiviCorret(t *testing.T) {
	teste := divisao(90, 3)
	resultado := 0
	if teste != resultado {
		t.Error("Valor esperado :", resultado, "Valor retornado:", teste)
	}
}
func TestDiviIncorret(t *testing.T) {
	teste := divisao(90, 3)
	resultado := 44
	if teste != resultado {
		t.Error("Valor esperado :", resultado, "Valor retornado:", teste)
	}
}
func TestSubCorrect(t *testing.T) {
	teste := subtra(480, 22, 15)
	resultado := -517

	if teste != resultado {
		t.Error("Valor esperado :", resultado, "Valor retornado:", teste)
	}
}
func TestSubIncorrect(t *testing.T) {
	teste := subtra(480, 22, 15)
	resultado := 1000

	if teste != resultado {
		t.Error("Valor esperado :", resultado, "Valor retornado:", teste)
	}
}
