# Concurrent rule simulation in

> It takes only a few simple rules to go from chaos to order

To run the program, go to the root folder and type

```
go run .
```

## Options

Open the `main.go` file and set your own parameters for

1. `Count` - how many particles are simulated
2. `MaxSpeed` - how fast can they go. Putting this too high results in particles leaving the viewport
3. `MaxDistance` - how far can particles be to still have effect on each other
4. `SlowingCoefficient` - what multiplier do the particles get slowed down by on each iteration. This helps stabilize the system.
5. `Height` and `Width` - the sizes of the viewport

## Rules

The rules are randomized on each run. If you want your own custom rules, open the `setup.go` file, remove the `randG()` calls and put your own `float64` numbers there.
