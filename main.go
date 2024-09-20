package main

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"time"

	"golang.org/x/exp/rand"
)

func main() {
	fmt.Println(ascii_archive())

}

func ascii_archive() string {
	// Seed the random number generator
	rand.Seed(uint64(time.Now().UnixNano()))

	//array of strings
	art := []string{
		`

	1. Cute Cat brings you fortune
	 /\_/\  
	( o.o ) 
	 > ^ < 
`,
		`
Tiny Bear wishes you good luck

ʕ•ᴥ•ʔ
`,
		`
Bunny says relax

(\(\ 
( -.-)
o_(\")(\")
`,
		`
Koala says you are cute

  (\/)
 (o.o)
  (> <)
`,
		`
Cute Bird brings you food

  \\
  (o>
\\_//)
 \_/_)
  _|_
`,
		`
Smiling Frog tells you to be immune to failure like his skin

   @..@
  (----)
 ( >__< )
 ^^  ^^
`,
		`
Sleepy Owl grants you crazy amounts of luck

  ,___,
  [O.o]
  /)  )
 -"-"-
`,
		`
Chubby Pig hopes you grow chubby like him from success

  ^  ^
 ( 'oo' )
  (    )
  ^^ ^^
`,
		`
Jumping Bunny wishes you luck

	(\(\ 
	(='.')
   c(_(\")(\")
`,
		`
Tiny Bird tells you to be proud of yourself
   
	 \\
	 (o>
   \\_//)
	\_/_)
	 _|_
   `,
		`
cute ducky wants you to feed her

   <(o )___
	(  ._> /
	 '---' 
`,
		`
Sleepy Bird suggests you sleep it off

   zZz
  ~(_o)
   / \
`,
		` 
Singing Bird sings your glory

 ♪ ♫ ♪
 (o>
 / ) )
/  | |
`,
		`
♡ ∩_∩ 
  („• ֊ •„)♡
|￣U U￣￣￣￣￣￣￣￣￣|
|        I love you    |   
￣￣￣￣￣￣￣￣￣￣￣￣
`,
		`
    /)  /)
ପ(˶•-•˶)ଓ ♡
   /づ  づ
`,
		`
bunny wished the other one were you                 




                                                        ★ ⁺.

                                        (\_(\    /)_/)
                                        (    )  (     )
                                       /    |    |    \
                                      ( O   |    |   O )
`,
		`
cute cat brings you money
  ^~^  ,
 ("Y") )
 /   \/ 
(\|||/)
`,
	}
	randomart := art[rand.Intn(len(art))]
	return randomart

}

// // Map grayscale brightness to ASCII characters
// func brightnessToChar(brightness float64) string {
// 	ascii := " .:-=+*#%@"
// 	index := int(brightness * float64(len(ascii)-1))
// 	return string(ascii[index])
// }

//	func ascii_cats() {
//		fmt.Println("here they are")
//		fmt.Println("  ^~^  ,")
//		fmt.Println(" ('Y') )")
//		fmt.Println(" /   \\/ ")
//		fmt.Println("(\\|||/)")
//	}

// func cats_from_cataas() {
// 	resp, err := http.Get("https://cataas.com/cat")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()
// 	img, _, err := image.Decode(resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// Resize and convert image to grayscale
// 	height := 100
// 	width := 100 // Adjust width to fit your terminal
// 	img = imaging.Resize(img, width, height, imaging.Lanczos)
// 	grayImg := imaging.Grayscale(img)

// 	// Create ASCII art by mapping brightness to characters
// 	var asciiArt strings.Builder
// 	for y := 0; y < grayImg.Bounds().Dy(); y++ {
// 		for x := 0; x < grayImg.Bounds().Dx(); x++ {
// 			grayColor := grayImg.At(x, y)
// 			r, g, b, _ := grayColor.RGBA()
// 			// Calculate brightness: average of r, g, and b (scaled to 0-255 range)
// 			brightness := float64((r+g+b)/3>>8) / 255.0
// 			asciiArt.WriteString(brightnessToChar(brightness))
// 		}
// 		asciiArt.WriteString("\n")
// 	}
// 	fmt.Println(rand.New(rand.NewSource(9)))
// 	// Print ASCII art
// 	fmt.Println(asciiArt.String())

// }
