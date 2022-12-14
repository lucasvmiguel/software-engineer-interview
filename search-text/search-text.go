package searchtext

import (
	"regexp"
	"sort"
	"strings"
)

type SearchText struct {
	forwardIndex  map[int][]string
	invertedIndex map[string][]int
}

type Doc struct {
	ID    int
	Score float32
}

func New() *SearchText {
	return &SearchText{
		forwardIndex:  make(map[int][]string),
		invertedIndex: make(map[string][]int),
	}
}

func (s *SearchText) Feed(id int, text string) error {
	keywords, err := s.tokenize(text)
	if err != nil {
		return err
	}

	s.forwardIndex[id] = keywords

	for _, keyword := range keywords {
		_, ok := s.invertedIndex[keyword]
		if ok {
			s.invertedIndex[keyword] = append(s.invertedIndex[keyword], id)
		} else {
			s.invertedIndex[keyword] = []int{id}
		}
	}

	return nil
}

func (s *SearchText) Search(text string) ([]Doc, error) {
	keywords, err := s.tokenize(text)
	if err != nil {
		return nil, err
	}

	result := []Doc{}
	for _, keyword := range keywords {
		docIDs := s.invertedIndex[keyword]
		for _, docID := range docIDs {
			result = append(result, s.tfi(docID, keyword))
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Score > result[j].Score
	})

	return result, nil
}

func (s *SearchText) tfi(id int, keyword string) Doc {
	var appearences int
	docKeywords := s.forwardIndex[id]
	for _, docKeyword := range docKeywords {
		if docKeyword == keyword {
			appearences++
		}
	}

	return Doc{
		ID:    id,
		Score: float32(appearences) / float32(len(docKeywords)),
	}
}

// TODO
// - proper tokenize
func (s *SearchText) tokenize(text string) ([]string, error) {
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		return nil, err
	}
	str := re.ReplaceAllString(strings.ToLower(text), " ")
	return strings.Split(str, " "), nil
}
