package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	assets := []string{
    "texture_grass.png",
    "sound_explosion.wav",
    "model_player.obj",
    "shader_basic.vert",
    "font_main.ttf",
    "texture_skybox.png",
    "sound_background.mp3",
    "model_enemy.obj",
    "shader_lighting.frag",
    "config_game.json",
}


	for _, asset := range assets {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()

			fmt.Printf("Loading %s...\n", name)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Printf("Finished loading %s\n", name)
		}(asset)
	}
	wg.Wait()
	fmt.Println("All game assets loaded. Starting game...")
}
