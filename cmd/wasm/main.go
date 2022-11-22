package main

import (
	"fmt"
	"math/rand"
	"syscall/js"
)

type sessionType struct {
	CardIndex int
}

var session sessionType

var sessionCards []card

func newSession() js.Func {
	jsonfunc := js.FuncOf(func(this js.Value, args []js.Value) any {

		jsDoc := js.Global().Get("document")
		if !jsDoc.Truthy() {
			result := map[string]any{
				"error": "Unable to get document object",
			}
			return result
		}
		outputArea := jsDoc.Call("getElementById", "cardsRemaining")
		if !outputArea.Truthy() {
			result := map[string]any{
				"error": "Unable to get output area",
			}
			return result
		}
		outputArea.Set("textContent", len(cards))

		characterArea := jsDoc.Call("getElementById", "character")
		if !characterArea.Truthy() {
			result := map[string]any{
				"error": "Unable to get output area",
			}
			return result
		}

		pinyinArea := jsDoc.Call("getElementById", "pinyin")
		if !pinyinArea.Truthy() {
			result := map[string]any{
				"error": "Unable to get pinyin area",
			}
			return result
		}

		meaningArea := jsDoc.Call("getElementById", "meaning")
		if !meaningArea.Truthy() {
			result := map[string]any{
				"error": "Unable to get meaning area",
			}
			return result
		}

		sessionCards = make([]card, len(cards))
		copy(sessionCards, cards)
		rand.Shuffle(len(sessionCards), func(i, j int) {
			tc := sessionCards[i]
			sessionCards[i] = sessionCards[j]
			sessionCards[j] = tc
 		})

		session.CardIndex = 0
		characterArea.Set("textContent", sessionCards[session.CardIndex].character)
		pinyinArea.Set("textContent", sessionCards[session.CardIndex].pinyin)
		meaningArea.Set("textContent", sessionCards[session.CardIndex].meaning)

		return nil
	})
	return jsonfunc
}

func answerCard(difficulty int) js.Func {
	jsonfunc := js.FuncOf(func(this js.Value, args []js.Value) any {

		jsDoc := js.Global().Get("document")
		if !jsDoc.Truthy() {
			result := map[string]any{
				"error": "Unable to get document object",
			}
			return result
		}

		if 0 != len(sessionCards) {

			switch difficulty {
			case 0:
				// And remove the card from the session
				sessionCards = append(sessionCards[:session.CardIndex], sessionCards[session.CardIndex + 1:]...)
			case 1:
			case 2:
			default:
				panic("bad difficulty")
			}

			outputArea := jsDoc.Call("getElementById", "cardsRemaining")
			if !outputArea.Truthy() {
				result := map[string]any{
					"error": "Unable to get output area",
				}
				return result
			}
			outputArea.Set("textContent", len(sessionCards))

			if 0 == len(sessionCards) {
				// Completed

			} else {

				characterArea := jsDoc.Call("getElementById", "character")
				if !characterArea.Truthy() {
					result := map[string]any{
						"error": "Unable to get output area",
					}
					return result
				}

				pinyinArea := jsDoc.Call("getElementById", "pinyin")
				if !pinyinArea.Truthy() {
					result := map[string]any{
						"error": "Unable to get pinyin area",
					}
					return result
				}

				meaningArea := jsDoc.Call("getElementById", "meaning")
				if !meaningArea.Truthy() {
					result := map[string]any{
						"error": "Unable to get meaning area",
					}
					return result
				}

				session.CardIndex += 1
				session.CardIndex = session.CardIndex % len(sessionCards)
				characterArea.Set("textContent", sessionCards[session.CardIndex].character)
				pinyinArea.Set("textContent", sessionCards[session.CardIndex].pinyin)
				meaningArea.Set("textContent", sessionCards[session.CardIndex].meaning)
			}
		}

		return nil
	})
	return jsonfunc
}

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("newSession", newSession())
	js.Global().Set("easyCard", answerCard(0))
	js.Global().Set("mediumCard", answerCard(1))
	js.Global().Set("hardCard", answerCard(2))
	<-make(chan bool)
}