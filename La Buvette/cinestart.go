package rpg

import (
	"bufio"
	"fmt"
	"os"
)

var nomValide string

func Cinestart(p *Player) {
	ClearScreen()
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                  ▒▒▒█▓                                           ")
	fmt.Println("                                                ▒▒▒▓█▓▓▓▓▓▒                                       ")
	fmt.Println("                                              ▒▒▒▓▓▓▒▒▒▒▓▓▓▓▓▓                                    ")
	fmt.Println("                                             ░▒░▒▓▒░░▒▒▒▒▒▓▓▓▓█▓▒                                 ")
	fmt.Println("                                            ▒▒▒▓▒░░░░░░░▒▒▒▓▓▓▓▓▓▓▓▒                              ")
	fmt.Println("                                        ▓▓▓▓▒▒▓▒▒░░░▒░░░▒▒▒▓▓▓▓▓███▓▓                             ")
	fmt.Println("            ▒▓▓███▓▒                  ▒▓▓▓▓▓▒▓▓▒▒░░░▒▒▒▓▒▒▓▓▓▓▓▓█████▒                            ")
	fmt.Println("           ▒▓▓▓▓█▓▓▓▒▒                ▓▓▓▓▓▓▓▓▒▒▒░░▒▒▓▓▒▒▓█▓▓▓▓▓▓███▓▓▓                           ")
	fmt.Println("          ▓▓▓▓▓█▓▓▓▓▓▓▒▒             ▓▓▓▓▓▒▒▒▒▒▒░░░▒▒▒▒▒▒▓█████████▓▓▓▓▓                          ")
	fmt.Println("       ▓▓▓▓▓▓▓▓█▓▓▓▓▓▓█▓▓            ▓█▓▒▓▓▓▓▒▒▒░░▒▒▓▒▒▒▓▓▓█▓▓██████▓▓██▓▒▒                       ")
	fmt.Println("     ▒████▓▒▒▓▓▓▓▓▓▓▓▓███▒         ▒▒███▓▓▓▓▓▓▓▓▓▒░▒▓▒▒░▓▓▓▓▓█████████████▓▓▓▒                    ")
	fmt.Println("     ▒█████▓▒██▓▓▓▓▓▓▓████▓▓▒       ▒█████████████▓█████▓▓▓▓▓▓█████████▓▓█▓███▓▓                  ")
	fmt.Println("     ▒██████▓▓▓▓▓▓▒████████▓▓▒ ▒     ███▓▒▒░▒▒▒▓▓▓████████▓▓▓▓▓▓▓████▓▒░▒▓█████▓▒                 ")
	fmt.Println("      ▒██████▓▓▓██████████▓▓▓▒▒    ▒▓▓▓▓█▓▓▓▓▓▒▒▒▓▒▓█▓███████▓▓█████▓▒░░░░███████▒                ")
	fmt.Println("       ▓████████████████▓▓██▓   ▒▓▓▓▓▓▓▓████▓▓▓▒▒▒▓▓▓▓▓█▓█████▓▓████▒░░░░░███████▒                ")
	fmt.Println("        ▒▓████▓▓████████████▒    ▒██████▓███████▓▓▓▓▓▓▓▓▓▓▓██████████▓▒▒▓▓███▓███▒                ")
	fmt.Println("          ▒▒▒▒  █████████████  ▓ ▓▓███▓▓▓▓▓▓▓▓██████████▓▓█▓▓████████▓▓▓▓██▓ ▓████▓               ")
	fmt.Println("                ▓██████▓▓██▓█████▓▓▓▓▓▒▓█▓▓▒▒▒▓▓▓▓███▓▓███▓▓▓▓░░░▓▓▓▓▓▓█▓▓   ▓███▓▓               ")
	fmt.Println("                ▓▓████▓██▓▓██▓▓██▓▓▓▓▒▓▓▓▓▓▒▓▓▓▓▒▒▓▓▓▓████▓▓▓▓▒░░▒▒▒▓▓█▓▒     ▒▒▒                 ")
	fmt.Println("                 ▓████▓████████▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓███▓▓▓▓▓▓░▒▒░░░▒▓▓▓▓▓▒                       ")
	fmt.Println("                   ███████▓▓█████████▓▓▓▒▓▓▓▓▓▓▓▓█▓▓▓█████▓▓▓▓▓▓▒▒▒▓███████▒                      ")
	fmt.Println("                   ▓██████▓▓████▓▓██▓█▓▓▓▓▓▓▓▓▓▓▓▓███▓█████▓▓▓███▓▓▓▓▓████▓▒                      ")
	fmt.Println("                   ▓▓▓ ▓███▓▓███▓▓█▓█████▓▓▓███▓▓▓███████████▓▓▓█▓▓▓▓▓████▓▒                      ")
	fmt.Println("                           ▒████████▓▓██▓▓▓███████▓▓▓███▓██████████▓▓█████▓▒                      ")
	fmt.Println("                         ▒▓█▓▓▓████████▓▓▓▓████████▓▓████████████▓▓▓▓▓▓▓▓▓                        ")
	fmt.Println("                         ███▓█████████████████████▓▓▓████████▓██▓▒▓▓▓▒▒▓▓▓                        ")
	fmt.Println("                          ▓██▓▓████▓▓████████████████████████▓██▓█▓▓▓▓█████▒                      ")
	fmt.Println("                          ▓█▒  █████▓▓██████████████▓███████▓▓▓▓▓▓▓▓▒▓█████▒                      ")
	fmt.Println("                               ▒████████████████████████████▓▓▓▓▓▓██▓█████▓▒                      ")
	fmt.Println("                               █████████████▓▓█▓████▓▒▓█▓████▓▓▓▓▓▓▓███▓▓▓                        ")
	fmt.Println("                              ▒████████████████▒▓███████▓████████████▓▒                           ")
	fmt.Println("                             ▓███████████▓▓▓▓▓▓▒▓██████▓▓▓▓█████████▓▒                            ")
	fmt.Println("                             ▓████████▓▓▓██████▓▓████▓▓▓▓▓▓█▓▒▓▒▒                                 ")
	fmt.Println("                             ▒▒████████▓▓▓▓▓▓██▓▓█▓▓▓▓▓██████▓                                    ")
	fmt.Println("                                ▓█████████▓▓▓▓█▓▓▓▓▓████████▓                                     ")
	fmt.Println("                              ▒████████████▓████████████████▒                                     ")
	fmt.Println("                            ▒▓████████████▓▓▓▓▒▒▓████▓▓▒▒▓██▒                                     ")
	fmt.Println("                            ▒▓▓▓▓▓▓▓█████▒▒░░░░░▒████▓▓▓▓▓██▒                                     ")
	fmt.Println("                                                 ▓▓▓▓███████▓▒                                    ")
	fmt.Println("                                                                                                  ")
	fmt.Println("Gurdil : Hey ! Salut ! Bienvenue à ma soirée ! Viens entre ! Eh d'ailleurs comment tu t'appelles ?!\n")

	fmt.Print("C'est quoi ton nom ? : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	player.Name = scanner.Text()
	scanner.Scan()

	ClearScreen()
	fmt.Println("Gurdil : Eh beehh bienvenue alors ! Amuses toi bien !\n")
	fmt.Println(" ")
	fmt.Println("                                                                                                     ")
	fmt.Println("                                            ▓▒           ▓▓▒ ▒ ▓                                     ")
	fmt.Println("                                         ▒▒▓▒▒ ▓▓▓▒▓▒    ▒▒▓▓▒ ▓▓▓                                   ")
	fmt.Println("                  ▒▒▒▒▒                ▒▒▓▓▓▓▒▒▓▓▓▓▓▓▒▒▒ ▒ ▒▒▓▒██▓▓                                  ")
	fmt.Println("                  ▓▓▓▓▓▒▒▒▒            ▓███████▓▓▓▓▒▓▓▓▒▒▒▒▒▒▓▓██▓▓▒▒                                ")
	fmt.Println("              ▒▒▒▒▓▓▓▓▒ ▓▓▓▒▒▒          ████████▓▓▓▓▓▓▓█▓▓▓▓▓▓▓██▓▓▓▓▒                               ")
	fmt.Println("              ▓█▓▓█▓▒▒▒  ▒  ▓▓▒▒▒▒▒     ████████████▓▓▓▒▓█▓▓▓▓▓██▓▓▓▓▓▒                              ")
	fmt.Println("             ▒███▓█▓▓▓▒▒ ▒▒▒▒ ▓▓█▓▓▒    ▓▓▓▓█▓▓▓▓▓████▓▓▓▓▓█▓▓▓▓▓▓▓▓▓▓▓▒                             ")
	fmt.Println("             ▓█▓██▓▓▓▓▓▒  ▓▓▒█▓▓█▓█▓    ▓▓▓▓▓▓▓████▓████▓▓▓███▓▓▓▓▓▓▓▓▓▓▒                            ")
	fmt.Println("             ▒█▓▓█████▓▓▓█▒▓█▓███▓▒      ▓▓▓█████▓▓▓██▓██▓████▓▓▓▓▓▓▓▓▓▒                             ")
	fmt.Println("              ▒▒▓███████▓▓▓▓██████▒      ▒▒▒▒▒▓▓▒▓▓▓▓▓▓█▓▓▓██▓█▓▓▓▓▓▓▓▒                              ")
	fmt.Println("                ▒▒▓▓▓▓██▓▓█▓▓██▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▒▒▒▒▒▓█████▓▓▓▓▓▓▓▓▓▓▒                               ")
	fmt.Println("                ▒ ▒▒▒▒▓███████▓▒▒▓▓▓▓███████▓▓▓▓▓▓▓▓▓▓█▓▓██████▓██▓▒▒                                ")
	fmt.Println("                ▒▒▒▒▒▓▓██████▓▒  ▒▓▓▓███████▓▓▓▓▓▓▓▓██▓▓▓█████████▓▒                                 ")
	fmt.Println("              ▒▒▒▓▓▓▓▓▓▓▓▓▓▓▒    ▒▓██████████▓▓▓██████▓▓▓▓██▓▓▓█▓▓▓                                  ")
	fmt.Println("             ▒▓▓▓▓▓▓████▓▓▒  ▒▒▒▒▓▓█████▓███████████▓▓▓▓▓▓▓███████▓▒                                 ")
	fmt.Println("             ▓▓██▓▓▓▓██▓▓▓   ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓███████▓▓▓▓▓▓▓████████▓▒▒  ▒▒                            ")
	fmt.Println("             ▓▓▓▓▓▓███▓▓▓▒▒  ▓▓▓██▓▓▓▓▓▓▓▓▓▓▓▓██▓▓▓█▓▓▓▓▓████████▓▓▓▓▒▒▓▓▒▒                          ")
	fmt.Println("               ▓▓█████▓▓▓▓▓██▓████▓▓▓▓▓▓▓▓▓▓▓▓▓█▓▓▓▓████████████▓▓▓▓▓▓▒ ▓█▓▒                         ")
	fmt.Println("               ▓▓████▓▓██▓█████▓▓██▓▓▓▓██▓▓▓▓▓▓██▓▓▓▓▓███████▓█▓▓▓▓▓▓▒█████▓                         ")
	fmt.Println("               ▓██████▓██▓██▓█▓▓▓▓▓▓▓▓▓████▓▓▓████████████████▓▓▓▓▓▓█▓███▓▒▒                         ")
	fmt.Println("              ▓███▓▓▓▓▓▓▓▓▓███▓▓▓▓▓▓▓▓▓▓▓██████████▓████████▓▓▓▓▓▓▓▓███▓▓▓▒                          ")
	fmt.Println("              ▓▓▓▓▓█████▓▓▓███▓▓▓▓▓▓▓▓▓▓▓███████▓▓███████████▓▓▓▓▓███▓▓▓▓▓▒                          ")
	fmt.Println("              ▓▓▓▓▓██████▓▓████▓▓▓▓▓▓▓▓▓▓██████▓▓▓███████████▓▓▓▓▓██▓▓▒▒▒▓▒   ▒                      ")
	fmt.Println("              ▓▓▓▓██▓▓█████████████████▓██████▓▓▓█████████▓██▓▓▓▓▓█▓▓▒▒ ▒▓▒   ▓▒                     ")
	fmt.Println("              ▓▓▓▓██▓▓█▓▓███████████████████▓██▓███████████▓▓▓▓████▓▒▒▒▒▒     ▓▓▒                    ")
	fmt.Println("                  ▓▓▓▓▓▓▓█████▓▓▓███████▓▓████████████████▓▓▓▓████▓▓▓▓▓▓▒▒▒▒▒ ▒▓▓▒                   ")
	fmt.Println("                     ▓▓▓▓▓█████▓▓▓██████▓█████████████████▓▓▓▓▓▓▓▓▒ ▒▓▓▓▓▓▓▓▒▒▓▓▓▓▒                  ")
	fmt.Println("                          ▓████▓▓███████████████████████████████▓▓   ▓▓▓▓▓▓▒ ▓▓▓▓▓▓▓▓▒               ")
	fmt.Println("                          ▒▓███████████████████▓██▓████████████▓▓▓   ▓▓█▓▒███▓▓██▓██▓▓               ")
	fmt.Println("                           ▒▓▓▓▓▓▓██████████▓█████████▓███████▓▓▒▒    ███▓▓▓▓▒▓██████▓▒              ")
	fmt.Println("                            ▓▓██████████▓▓▓▓▒▓████████████████▓▒     ▓▓▓█▓▓▓▓▒▓███▓▓▓▒               ")
	fmt.Println("                          ▒▓▓▓█████████▓████▓▓▓██▓▓▓▓██████████     ▒▓▓▓▓█▓▓▓▓▓█▓█▓▓▒                ")
	fmt.Println("                          ▓▓▓▓████████▓▓█████▓▓██▓▓▓▓██████████▓▒▒▒▒▓▓▓▓▓█▓▓▓▓▓█▓█▓▓                 ")
	fmt.Println("                          ▒▓█▓▓███████▓▓█████▓▓██▓▓▓▓██████████▓▓▓▓▓▓▓██▓█▓▓▓▓██▓▓▓                  ")
	fmt.Println("                           ▓▓▓▓▓███████████████████████████████▓▓▓█▓▓▓██▓▓▓▓▓█▓▓▓▓▒                  ")
	fmt.Println("                                ████████▓███████▓█████████████▓▓▓▓████▓▓▓▓▓▓▓████▒▒▒                 ")
	fmt.Println("                                ████▓███▓▓██████▓█████████████▓▓▓▓███████▓▓▓█▓███▒▓▒▒                ")
	fmt.Println("                                 ███████▓███████▓▓██████▓▓███████▓███████▓▓▓█▓█▓▓▓▓▓▓                ")
	fmt.Println("                              ▓▓▓▓██████▓▓▓█████▓▓▓▓ ▒▓▓███████████▓▓▓▒▓█▓▓▓▓▒▒▒▒▓▒▒▒                ")
	fmt.Println("                             ▓▓▓▓▒▓█████▓███████▒     ▓████▓▓▓▓▓█████▓▒▒▓▓██▓▓▒                      ")
	fmt.Println("                            ▓████▓████████████▓▓      ▓██████▓████████▒ ▓▓▒▒▒▓▒                      ")
	fmt.Println("                            ▓██▓▓▓▓▓█▓▓▓▓▓▓▓▓▓▒▒      ▒▓▓▓▓██████████▓  ▒▒   ▒                       ")
	fmt.Println("                            ▒▓▓▒▒▒▒▒▓▒▒▒▒▒▒▒▒▒         ▒▒▒▒▓▓▓▓▓▓▓▓▓▓▒                               ")
	fmt.Println("                                                                                                     ")
	fmt.Println("**Gloup Gloup**")
	fmt.Println(" Passpatou : Ehhh salut ! Va te chercher une bière et viens nous rejoindre !")

	scanner.Scan()

	ClearScreen()
	fmt.Println("")
	fmt.Println("                                                                                                     ")
	fmt.Println("                                           ▓                                                         ")
	fmt.Println("                                          ▓░▒                                                        ")
	fmt.Println("                                       ▓▒▓▓░▒▒▒▓                                                     ")
	fmt.Println("                                    ▓▒░░▒▓▓░▒▓░░░▒▓                                                  ")
	fmt.Println("                                  ▓▒▒░░░▓▓▓░▒▓░░░░░░▓                                                ")
	fmt.Println("                                 ▓▓▒▒░░▒▓▓▓▒▒▓░░░░░░░▒█                                              ")
	fmt.Println("                                ▓▓▓▒▓▒▒▒▓▓▓▓▒▓▒▒▒░░░░░▓              ▓▓█▓▒▓███▓▓▒▓▓                  ")
	fmt.Println("                               ▓▓▓▓▓▓▓▒▓█▓▓▓▓▓▒▒▒▒▓▒▒░▒▓          ▓▒░░░░░░░░▒▒░░░░░▒▓                ")
	fmt.Println("                              ▓▓▓▓▓▓▓▒▓▓▓▓▓▒▒▒▒▒▒▒▒▒▒▒▓▒▒       ▒░░░░░░░░░░░░░░░░░░░░░▓              ")
	fmt.Println("                              ▓▓▓▓▓▓▓▓▓▓▓▓▒░▒▓▓▓▓▓▓▓▒▒▒▒▒       ▒░▒▒░░░░░░░░░░░░░░░░░░▒              ")
	fmt.Println("                             ▓▓████▓▒░▒▒▓█▒░▓▓▒░░░▒▓██▓▓▓    ▓▒▓▓▓▓▓▒▒▒░░░░░░░░░░░░░░░░              ")
	fmt.Println("                            ▓▓█████▓▓▒▒▒▓▓███▓▓▒▒▒▒▓██▓██▓  ▓▒▒  ▒▓▒▒▒▒░░░░░░░░░░░░░░░░              ")
	fmt.Println("                           █▓███████▓▓▓▓▓▓▓█▓▓▓▓▓▓▓▓█▓▓██▓  ▓▒    ▓▒▓▓▓▓░░▒░░░░░░░░░░░▒              ")
	fmt.Println("                         █▓███████▓▓▓██████▓█████▓▓▓▓█▓█▓███▓▒   ▓▓▓▒▓░▓▒▒▓▓▓░▒░▒▓▓▓▒░▒              ")
	fmt.Println("                    ▓▓▒▒▓█▓▓██████████▓▓██▓▓▓████████▓▓█▓██▓▓▒   ▓▓▓▒▓▒▓▒▒▓▒▓▓▓▒▓▒▒▓▒▓▒              ")
	fmt.Println("                  ▓▒░░▒▓▓███████▓███▓██▓█▓▓▓█▓▓▓█▓▓██▓██▓▓▓███▓▓ ▓▓▓▒▓▒▓▒▒▓▒▒▓▒▒▓▒▒▓▒▓▒              ")
	fmt.Println("                 ▓▓▓▒▒▒▓████████▓▓████▓▓██▓▓█▓▓▓█▓▓█▓▓▓█▓▓█▓▒▓▓▒  ▓▓▒▓▒▓▒▒▓▒▒▓▒▒▓▒▒▓▒▓▒              ")
	fmt.Println("                 ██▓▒▒▒▓▓▓█████████▓▓██████▓██▓██▓▓█▓▓▓████▓▒▓▓▒▓ ▓▓▒▓▒▓▒▒▓▒▒▓▒▒▓▒▒▓▒▓▒              ")
	fmt.Println("                ▓▒▒▓▓▓▓▓▒▓████▓███████▓▓██▓▓█▓▓█▓▓▓███▓██▓█▒▒▓█▒▓▓▓▓▒▓▒▓▒▒▓▒▒▓▒▒▓▒▒▓▒█░              ")
	fmt.Println("              ▓▒▒░░░▒▓█▓▓▓█▓█████████████▓▓▓██▓▓██▓█▓▓█████▓▒▓█▒▓▓▓▓▒▓▒▓▒▒▓▒▓▓▒▒▓▓▒▓▓▓▒              ")
	fmt.Println("             ▓▒▒▒░░░░▒▓▓█▓█▓███████▓████▓▓▓▓█▓▓▓████▓▓█████▓▒▓█▒▓▓▓▒▓▒▒▓▒▒▒▒▒▓▒░▒▒▒▓▒▒▒              ")
	fmt.Println("            █▒▓▒░░░░░░░▒▓▓█▓▓▓▓██████████▓███▓██████████▓█▓▒▒▓▓▓▓█▓▓▓▓▓▓▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒             ")
	fmt.Println("            ▓▒▒░░░░░░░▒▒█▓██▓▒▓███████████▓▓██████████▓▓▒█▓▒░░▒▒▒▒▓▓▓▓▓▓▒▒▒▒░▒░░░░▒░░░▒░▒            ")
	fmt.Println("          ▓▓▓▒▒▒▒▒▒▒▒▒▒▓█▓█▓▓▓▒▓▓▓▓▓▒▓█████▓▓▓██████▓▒▒░▒▒██▓▓▓▒▒▓ ▓▓█▓█▓▓▓▓▓▓▓▓▓▓▓▓▓▓               ")
	fmt.Println("       ▓▓▓▓▓▒▒▒▒▒▒▒▒▒▒▓▓▓███▓▓▓▓▒▒▒▒▒▒▓████▓▓▓███▓▒▓▒▒░░░▒▓█▓▓       ▓▓▒▒▒▒░░▒▒▒▒▓▓▓                 ")
	fmt.Println("       ▓▓▓▓▓▒▒▒░░░░▒▓▓ ▓▓▓▓▓▓▓▓▓▓▓▓▓▒▒▒▓▓███▓███▓▒▒▒▒▒▒▒▒▒▒▓▓▓                                       ")
	fmt.Println("       ▓▓▓▓▒▓▒▓▓▓▓▓▓▓▓  ▓▒▓▓▓▒▓█▓▓▓▓▓▓▓▓▒▒▓███▓▒▓▒▒▒▒▒▒▒▓▓▒▓█▒                                       ")
	fmt.Println("        ▓▒░░░░░░▒▓▓▓   ▓▓▓▓▓▒▒▓▒▓▓▓▓███▓▓▓▓▓▓▓▒▓▓▓▓▓██▓███▓▓▓                                        ")
	fmt.Println("       ▓▓▒░░░▒▒░░▓▓   ▓▓▒▒▒▓▓▓▓▓▒▒▒▓█▓▓▒██▓▓▓▓▒▓▓██▓▓██████▓▓▓                                       ")
	fmt.Println("       ▓▓▒▒▒▒░░▒▒▓▓  ▓▓▒▒▒▒▒▒▒▓▓▓▓███▓█▒▓█████▒▓▓▓▓▓▓█▓▓▓▓▓▓▓▓                                       ")
	fmt.Println("       ▒▒▒▒▒▒▒▒░▓▓   ▓▒▒▒▒▒▒▒▒▒▒▒████▓█▓▓▓▓▓▓▓▓█████▓▓▓▓▓▓▓▓▓▓                                       ")
	fmt.Println("       ▓▓ ▓▓▒▓▓       ▓▓▒▒░▒▒░░▓▓█▓████▓▓▓▓▓▓▓▓▓█▓▓█▓███████▓▓▓                                      ")
	fmt.Println("                       ▓▓▓▓▓▒▒▒▓▓▓▓▓▓████████████████▓▒▒▒░▓▓▒▒▒                                      ")
	fmt.Println("                         ▓▓▓▓▓██▓▓▒▒▓▓██████████████▓▒▓▒▒▒▒▒▓▓▒                                      ")
	fmt.Println("                         ▒▒▒▓█▓▓▓▓▒▒▒▒▓█ █ ██ ███████▓▓▓▒▓▒▒▓▓▓                                      ")
	fmt.Println("                         ▒▒▓▓▓▓▓▓▓███▓▓█       ██▓▓▓▓██████▓▓▓▒▒▒▓                                   ")
	fmt.Println("                       ▒▒▓▓▓▓▓▒▒▒▒▓▓▓█          █████████▓▓▒▒▒▒▒░▒░▒                                 ")
	fmt.Println("                      ▓▓▓▓▓▓▓▓▒▒▒▒▒▓▓▓█          ███████▓█▓▒▒▒▒░░░░░▒                                ")
	fmt.Println("                        ▓▓▓▓▓▓▓▓▒▒▒▓█▓           ████████▓▓▓▓▒▒▒▒▒▒░▒                                ")
	fmt.Println("                       ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓          ▓█▓█▓██▓▓▓▓▓▓▓▓▓▓▓▓▓                               ")
	fmt.Println("                                                                                                     ")
	fmt.Println("Barman : Salut ! Tiens, voici ta biere !")

	scanner.Scan()

	ClearScreen()
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                  ▒▒▒█▓                                           ")
	fmt.Println("                                                ▒▒▒▓█▓▓▓▓▓▒                                       ")
	fmt.Println("                                              ▒▒▒▓▓▓▒▒▒▒▓▓▓▓▓▓                                    ")
	fmt.Println("                                             ░▒░▒▓▒░░▒▒▒▒▒▓▓▓▓█▓▒                                 ")
	fmt.Println("                                            ▒▒▒▓▒░░░░░░░▒▒▒▓▓▓▓▓▓▓▓▒                              ")
	fmt.Println("                                        ▓▓▓▓▒▒▓▒▒░░░▒░░░▒▒▒▓▓▓▓▓███▓▓                             ")
	fmt.Println("            ▒▓▓███▓▒                  ▒▓▓▓▓▓▒▓▓▒▒░░░▒▒▒▓▒▒▓▓▓▓▓▓█████▒                            ")
	fmt.Println("           ▒▓▓▓▓█▓▓▓▒▒                ▓▓▓▓▓▓▓▓▒▒▒░░▒▒▓▓▒▒▓█▓▓▓▓▓▓███▓▓▓                           ")
	fmt.Println("          ▓▓▓▓▓█▓▓▓▓▓▓▒▒             ▓▓▓▓▓▒▒▒▒▒▒░░░▒▒▒▒▒▒▓█████████▓▓▓▓▓                          ")
	fmt.Println("       ▓▓▓▓▓▓▓▓█▓▓▓▓▓▓█▓▓            ▓█▓▒▓▓▓▓▒▒▒░░▒▒▓▒▒▒▓▓▓█▓▓██████▓▓██▓▒▒                       ")
	fmt.Println("     ▒████▓▒▒▓▓▓▓▓▓▓▓▓███▒         ▒▒███▓▓▓▓▓▓▓▓▓▒░▒▓▒▒░▓▓▓▓▓█████████████▓▓▓▒                    ")
	fmt.Println("     ▒█████▓▒██▓▓▓▓▓▓▓████▓▓▒       ▒█████████████▓█████▓▓▓▓▓▓█████████▓▓█▓███▓▓                  ")
	fmt.Println("     ▒██████▓▓▓▓▓▓▒████████▓▓▒ ▒     ███▓▒▒░▒▒▒▓▓▓████████▓▓▓▓▓▓▓████▓▒░▒▓█████▓▒                 ")
	fmt.Println("      ▒██████▓▓▓██████████▓▓▓▒▒    ▒▓▓▓▓█▓▓▓▓▓▒▒▒▓▒▓█▓███████▓▓█████▓▒░░░░███████▒                ")
	fmt.Println("       ▓████████████████▓▓██▓   ▒▓▓▓▓▓▓▓████▓▓▓▒▒▒▓▓▓▓▓█▓█████▓▓████▒░░░░░███████▒                ")
	fmt.Println("        ▒▓████▓▓████████████▒    ▒██████▓███████▓▓▓▓▓▓▓▓▓▓▓██████████▓▒▒▓▓███▓███▒                ")
	fmt.Println("          ▒▒▒▒  █████████████  ▓ ▓▓███▓▓▓▓▓▓▓▓██████████▓▓█▓▓████████▓▓▓▓██▓ ▓████▓               ")
	fmt.Println("                ▓██████▓▓██▓█████▓▓▓▓▓▒▓█▓▓▒▒▒▓▓▓▓███▓▓███▓▓▓▓░░░▓▓▓▓▓▓█▓▓   ▓███▓▓               ")
	fmt.Println("                ▓▓████▓██▓▓██▓▓██▓▓▓▓▒▓▓▓▓▓▒▓▓▓▓▒▒▓▓▓▓████▓▓▓▓▒░░▒▒▒▓▓█▓▒     ▒▒▒                 ")
	fmt.Println("                 ▓████▓████████▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓███▓▓▓▓▓▓░▒▒░░░▒▓▓▓▓▓▒                       ")
	fmt.Println("                   ███████▓▓█████████▓▓▓▒▓▓▓▓▓▓▓▓█▓▓▓█████▓▓▓▓▓▓▒▒▒▓███████▒                      ")
	fmt.Println("                   ▓██████▓▓████▓▓██▓█▓▓▓▓▓▓▓▓▓▓▓▓███▓█████▓▓▓███▓▓▓▓▓████▓▒                      ")
	fmt.Println("                   ▓▓▓ ▓███▓▓███▓▓█▓█████▓▓▓███▓▓▓███████████▓▓▓█▓▓▓▓▓████▓▒                      ")
	fmt.Println("                           ▒████████▓▓██▓▓▓███████▓▓▓███▓██████████▓▓█████▓▒                      ")
	fmt.Println("                         ▒▓█▓▓▓████████▓▓▓▓████████▓▓████████████▓▓▓▓▓▓▓▓▓                        ")
	fmt.Println("                         ███▓█████████████████████▓▓▓████████▓██▓▒▓▓▓▒▒▓▓▓                        ")
	fmt.Println("                          ▓██▓▓████▓▓████████████████████████▓██▓█▓▓▓▓█████▒                      ")
	fmt.Println("                          ▓█▒  █████▓▓██████████████▓███████▓▓▓▓▓▓▓▓▒▓█████▒                      ")
	fmt.Println("                               ▒████████████████████████████▓▓▓▓▓▓██▓█████▓▒                      ")
	fmt.Println("                               █████████████▓▓█▓████▓▒▓█▓████▓▓▓▓▓▓▓███▓▓▓                        ")
	fmt.Println("                              ▒████████████████▒▓███████▓████████████▓▒                           ")
	fmt.Println("                             ▓███████████▓▓▓▓▓▓▒▓██████▓▓▓▓█████████▓▒                            ")
	fmt.Println("                             ▓████████▓▓▓██████▓▓████▓▓▓▓▓▓█▓▒▓▒▒                                 ")
	fmt.Println("                             ▒▒████████▓▓▓▓▓▓██▓▓█▓▓▓▓▓██████▓                                    ")
	fmt.Println("                                ▓█████████▓▓▓▓█▓▓▓▓▓████████▓                                     ")
	fmt.Println("                              ▒████████████▓████████████████▒                                     ")
	fmt.Println("                            ▒▓████████████▓▓▓▓▒▒▓████▓▓▒▒▓██▒                                     ")
	fmt.Println("                            ▒▓▓▓▓▓▓▓█████▒▒░░░░░▒████▓▓▓▓▓██▒                                     ")
	fmt.Println("                                                 ▓▓▓▓███████▓▒                                    ")
	fmt.Println("                                                                                                  ")
	fmt.Println("EH toi ! Attention derriere to-")
	fmt.Println("*BBBBBBBOOOONNNNNKKKKKKKK*")

	scanner.Scan()

	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("                                                                                                  ")
	fmt.Println("\n\n\n... ...... ...... ....")

	scanner.Scan()

	fmt.Println("\n\n\nMais ... Ou suis-je ?.")

	Menu()
}
