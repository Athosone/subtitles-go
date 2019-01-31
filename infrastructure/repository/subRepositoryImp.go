package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	gourl "net/url"
	domain "subtitles/domain"
)

func failOnError(err error) {
	if err != nil {
		log.Fatalf("Failed: error:\n[%v]\n", err)
	}
}

type omDBSearchResultDTO struct {
	Search []struct {
		Title    string `json:"Title"`
		Year     string `json:"Year"`
		ImdbID   string `json:"imdbID"`
		ShowType string `json:"Type"`
	} `json:"Search"`
	TotalResults string `json:"totalResults"`
	Response     string `json:"Response"`
}

type SubRepositoryImp struct {
	client http.Client
}

func NewSubRepositoryImp() *SubRepositoryImp {
	lRet := SubRepositoryImp{}
	lRet.client = http.Client{}
	return &lRet
}

func (a *SubRepositoryImp) Find(episode domain.Episode) (domain.Subtitle, error) {
	url, error := gourl.Parse("http://www.omdbapi.com")
	failOnError(error)
	params := gourl.Values{}
	params.Add("apikey", "e6fe729a")
	params.Add("s", episode.ShowName)
	url.RawQuery = params.Encode()
	urlStr := url.String()
	resp, error := a.client.Get(url.String())
	failOnError(error)
	defer resp.Body.Close()
	res := omDBSearchResultDTO{}
	error = json.NewDecoder(resp.Body).Decode(&res)
	failOnError(error)
	resp.Body.Close()
	urlStr = "https://rest.opensubtitles.org/search/episode-%d/imdbid-%s/season-%d/sublanguageid-en"
	imdbId := res.Search[0].ImdbID
	urlStr = fmt.Sprintf(urlStr, episode.Number, imdbId, episode.Season)
	url, error = gourl.Parse(urlStr)
	failOnError(error)
	resp, error = a.client.Get(url.String())
	resOS := openSubtitleResultDTO{}
	error = json.NewDecoder(resp.Body).Decode(&resOS)
	failOnError(error)

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
