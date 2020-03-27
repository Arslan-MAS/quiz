package quiz
func ReturnValue(option int) string {
	switch option  {
	case 1:
        return "Covid19 cases in Pakistan :1,252"
     	case 2: 
        return "Covid19 cases in SouthKorea:9,332"
	case 3: 
        return "Covid19 cases in France:29,155"
 	case 4: 
        return " My Message : May Allah Protect Us All and Keep Us All Safe "
	case 0:
        return " Exiting "
	default:
	return " wrong option"
	}	
}
