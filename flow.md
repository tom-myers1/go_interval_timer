* function flow for interval trainer

main()
  creates a temp config (required for menu)
  calls checkConfig() (sends nothing, receives nothing)    
  calls menu(Timer) (sends temp config)

checkConfig()
  looks for config file and creates if missing, errors if it cant

menu(Timer)
    receives Timer (type struct)
    creates menu with following options
      * press 1 to load from config file
                calls loadTimer() receives slice of timers
                calls selectTimer(timers) receives selected timer
                calls runTimer(selected)
      * press 2 to input settings
                calls userInput() returns timer
                calls runTimer(timer)
      * press 3 to save to config file
                calls saveTimer()
      * press 4 to delete a saved config
                calls deleteTimer()
      * press q to quit
                quits application

loadTimer()
    loads timers from config file
    returns slice of timers

selectTimer([]timers)
    expects slice of type struct timer
    returns timer

runTimer(timers)
      expects type struct timer
      runs timer

userInput()
      returns validated user inputted config
      calls runTimer()








current TODO's:
line 120
