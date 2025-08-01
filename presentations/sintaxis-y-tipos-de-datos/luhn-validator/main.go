package main

import "fmt"

// luhnCheck implementa el algoritmo de Luhn para validar números de tarjetas de crédito
func luhnCheck(cardNumber string) bool {
	sum := 0

	// Recorrer dígitos de derecha a izquierda
	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit := int(cardNumber[i] - '0')

		// Verificar que el carácter sea un dígito válido
		if digit < 0 || digit > 9 {
			return false
		}

		// Duplicar cada segundo dígito (contando desde la derecha)
		if (len(cardNumber)-i)%2 == 0 {
			digit *= 2

			// Si el resultado > 9, sumar sus dígitos
			if digit > 9 {
				digit = digit/10 + digit%10
			}
		}

		sum += digit
	}

	return sum%10 == 0
}

func main() {
	fmt.Println("🎪 Validador de Tarjetas de Crédito - Algoritmo de Luhn")
	fmt.Println("=======================================================")

	// Casos de prueba con diferentes tipos de tarjetas
	testCards := []string{
		"4532015112830366", // Visa válida
		"4532015112830367", // Visa inválida (último dígito cambiado)
		"5555555555554444", // Mastercard válida
		"5555555555554445", // Mastercard inválida
		"371449635398431",  // American Express válida
		"371449635398432",  // American Express inválida
		"6011111111111117", // Discover válida
		"1234567890123456", // Número aleatorio inválido
	}

	fmt.Println("\n🔍 Probando números de tarjetas:")
	fmt.Println()

	for i, card := range testCards {
		isValid := luhnCheck(card)

		var cardType string
		switch {
		case len(card) == 15 && (card[:2] == "34" || card[:2] == "37"):
			cardType = "American Express"
		case len(card) == 16 && card[0] == '4':
			cardType = "Visa"
		case len(card) == 16 && (card[:2] >= "51" && card[:2] <= "55"):
			cardType = "Mastercard"
		case len(card) == 16 && card[:4] == "6011":
			cardType = "Discover"
		default:
			cardType = "Desconocido"
		}

		status := "❌ INVÁLIDA"
		if isValid {
			status = "✅ VÁLIDA  "
		}

		fmt.Printf("%d. %s | %s | %s\n",
			i+1, card, status, cardType)
	}

	fmt.Println()
	fmt.Println("💡 El algoritmo de Luhn detecta ~70% de errores de transcripción")
	fmt.Println("🔒 Usado por: Visa, Mastercard, American Express, Discover, etc.")
}
