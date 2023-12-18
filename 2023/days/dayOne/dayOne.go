type textNum int

type numRep struct {
	textNumber textNum
	intNumber int
	stringNumber string
}
const (
	one textNum = iota + 1
	two 
	three
	four
	five 
	six 
	seven 
	eight
	nine 
)

func (t textNum) String() string {
	return [...]string{"one","two", "three", "four", "five", "six", "seven", "eight", "nine"}[t-1]
}

func (t textNum) enumIndex() int {
	return	int(t)
}

func toTextNumInt( t int) textNum{
	return [...] textNum{one, two, three, four, five, six, seven, eight, nine}[t-1]
}

func toTextNum(str string) (textNum, error){
	var tNum textNum
	switch str {
	case "one":
		tNum = one
	case "two":
		tNum = two
	case "three":
		tNum = three
	case "four":
		tNum = four 
	case "five":
		tNum = five
	case "six":
		tNum = six
	case "seven":
		tNum = seven 
	case "eight":
		tNum = eight
	case "nine":
		tNum = nine
	default:
		return tNum, errors.New("not a number")
	}
	return tNum, nil
}



func parseLine(line string) []numRep{
	var numList []numRep

	for pos, char := range line {
		if unicode.IsDigit(char){
			num, _:= strconv.Atoi(string(char))
			tnum := toTextNumInt(num)
			numList = append(numList, numRep{tnum, tnum.enumIndex(), strconv.Itoa(tnum.enumIndex())})
			continue
		}

		var possibleNumStr string

		if( pos + 3 <= len(line)){
			possibleNumStr = line[pos: pos+3]
			tnum, err := toTextNum(possibleNumStr)
			if err == nil{
				numList = append(numList, numRep{tnum, tnum.enumIndex(), strconv.Itoa(tnum.enumIndex())})
				continue
			}
		}

		if pos + 4 <= len(line){
			possibleNumStr = line[pos: pos+4]
			tnum, err := toTextNum(possibleNumStr)
			if err == nil{
				numList = append(numList, numRep{tnum, tnum.enumIndex(), strconv.Itoa(tnum.enumIndex())})
				continue
			}
		}

		if pos + 5 <= len(line){
			possibleNumStr = line[pos: pos+5]
			tnum, err := toTextNum(possibleNumStr)
			if err == nil{
				numList = append(numList, numRep{tnum, tnum.enumIndex(), strconv.Itoa(tnum.enumIndex())})
				continue
			}
		}
	}
	return numList
}

func dayOne(){
	file, err := os.Open("./../inputs/inputOne.txt")

	if err != nil{
		fmt.Println("ERROR OPENING FILE")
	}

	scanner := bufio.NewScanner(file)

	total := 0
	p2Total := 0

	for scanner.Scan(){
		line := scanner.Text()
		var first rune = 'u'
		var last rune = 'u'
		numList := parseLine(line)

		if len(numList) > 0 {
			currNum  := "" + numList[0].stringNumber + numList[len(numList)-1].stringNumber
			fmt.Println(currNum)
			num, err := strconv.Atoi(currNum)
			if err == nil {
				p2Total += num
			}
		}

		for _, char := range line{
			if unicode.IsDigit(rune(char)){
				if first == 'u'{
					first = char
				}
				last = char
			}
			
		}
		currNum := "" + string(first) + string(last)
		num, err := strconv.Atoi(currNum)
		if err == nil {
			total += num
		}
	}
	fmt.Println(total)
	fmt.Println(p2Total)
}
