package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	domain "subtitles/domain"
)

func failOnError(err error) {
	if err != nil {
		log.Fatalf("Failed : %s error:\n[%v]\n", err)
	}
}

type omDBSearchResultDTO struct {
	search []struct {
		title    string `json:"Title"`
		year     string `json:"Year"`
		imdbID   string `json:"imdbID"`
		showType string `json:"Type"`
		poster   string `json:"Poster"`
	} `json:"Search"`
	totalResults string `json:"totalResults"`
	response     string `json:"Response"`
}

type SubRepositoryImp struct {
	client http.Client
}

func NewSubRepositoryImp() *SubRepositoryImp {
	lRet := SubRepositoryImp{}
	lRet.client = http.Client{}
	return &lRet
}

func (a *SubRepositoryImp) find(episode domain.Episode) (domain.Subtitle, error) {
	url := "http://www.omdbapi.com/?s=%s&apikey=e6fe729a"
	resp, err := a.client.Get(fmt.Sprintf(url, episode.ShowName))
	failOnError(err)
	defer resp.Body.Close()
	res := omDBSearchResultDTO{}
	err = json.NewDecoder(resp.Body).Decode(&res)
	failOnError(err)
	resp.Body.Close()
	url = "https://rest.opensubtitles.org/search/episode-%d/imdbid-%d/season-%d/sublanguageid-en"
	url = fmt.Sprintf(url, episode.Number, res.search[0].imdbID, episode.Season)
	resp, err = a.client.Get(url)
	resOS := openSubtitleResultDTO{}
	err = json.NewDecoder(resp.Body).Decode(&resOS)
	failOnError(err)

	subtitle := domain.Subtitle{}
	subtitle.Id = resOS[0].IDSubtitle
	subtitle.EpisodeName = resOS[0].MovieName
	return subtitle, nil
}

type openSubtitleResultDTO []struct {
	MatchedBy           string      `json:"MatchedBy"`
	IDSubMovieFile      string      `json:"IDSubMovieFile"`
	MovieHash           string      `json:"MovieHash"`
	MovieByteSize       string      `json:"MovieByteSize"`
	MovieTimeMS         string      `json:"MovieTimeMS"`
	IDSubtitleFile      string      `json:"IDSubtitleFile"`
	SubFileName         string      `json:"SubFileName"`
	SubActualCD         string      `json:"SubActualCD"`
	SubSize             string      `json:"SubSize"`
	SubHash             string      `json:"SubHash"`
	SubLastTS           string      `json:"SubLastTS"`
	SubTSGroup          string      `json:"SubTSGroup"`
	InfoReleaseGroup    string      `json:"InfoReleaseGroup"`
	InfoFormat          string      `json:"InfoFormat"`
	InfoOther           string      `json:"InfoOther"`
	IDSubtitle          string      `json:"IDSubtitle"`
	UserID              string      `json:"UserID"`
	SubLanguageID       string      `json:"SubLanguageID"`
	SubFormat           string      `json:"SubFormat"`
	SubSumCD            string      `json:"SubSumCD"`
	SubAuthorComment    string      `json:"SubAuthorComment"`
	SubAddDate          string      `json:"SubAddDate"`
	SubBad              string      `json:"SubBad"`
	SubRating           string      `json:"SubRating"`
	SubSumVotes         string      `json:"SubSumVotes"`
	SubDownloadsCnt     string      `json:"SubDownloadsCnt"`
	MovieReleaseName    string      `json:"MovieReleaseName"`
	MovieFPS            string      `json:"MovieFPS"`
	IDMovie             string      `json:"IDMovie"`
	IDMovieImdb         string      `json:"IDMovieImdb"`
	MovieName           string      `json:"MovieName"`
	MovieNameEng        interface{} `json:"MovieNameEng"`
	MovieYear           string      `json:"MovieYear"`
	MovieImdbRating     string      `json:"MovieImdbRating"`
	SubFeatured         string      `json:"SubFeatured"`
	UserNickName        string      `json:"UserNickName"`
	SubTranslator       string      `json:"SubTranslator"`
	ISO639              string      `json:"ISO639"`
	LanguageName        string      `json:"LanguageName"`
	SubComments         string      `json:"SubComments"`
	SubHearingImpaired  string      `json:"SubHearingImpaired"`
	UserRank            string      `json:"UserRank"`
	SeriesSeason        string      `json:"SeriesSeason"`
	SeriesEpisode       string      `json:"SeriesEpisode"`
	MovieKind           string      `json:"MovieKind"`
	SubHD               string      `json:"SubHD"`
	SeriesIMDBParent    string      `json:"SeriesIMDBParent"`
	SubEncoding         string      `json:"SubEncoding"`
	SubAutoTranslation  string      `json:"SubAutoTranslation"`
	SubForeignPartsOnly string      `json:"SubForeignPartsOnly"`
	SubFromTrusted      string      `json:"SubFromTrusted"`
	SubTSGroupHash      string      `json:"SubTSGroupHash"`
	SubDownloadLink     string      `json:"SubDownloadLink"`
	ZipDownloadLink     string      `json:"ZipDownloadLink"`
	SubtitlesLink       string      `json:"SubtitlesLink"`
	QueryNumber         string      `json:"QueryNumber"`
	QueryParameters     struct {
		Episode int    `json:"episode"`
		Season  int    `json:"season"`
		Imdbid  string `json:"imdbid"`
	} `json:"QueryParameters"`
	Score float64 `json:"Score"`
}
