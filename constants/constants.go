package constants

// The Kuali catalog URL must be manually updated. This one is up to date as of 2025 June 2
const CatalogUrl string = "https://uvic.kuali.co/api/v1/catalog/courses/68923109c72beab1a99e4c73"

// Kuali URL to get course information
const InformationUrl string = "https://uvic.kuali.co/api/v1/catalog/course/68923109c72beab1a99e4c73/%s"

const BaseUrl string = "https://banner.uvic.ca/StudentRegistrationSsb/ssb"
const CookieUrl string = BaseUrl + "/classSearch/classSearch?term=%s&txt_subject=CSUP&txt_courseNumber=000"
const SectionsUrl string = BaseUrl + "/searchResults/searchResults?txt_term=%s&pageMaxSize=10000&txt_subject=%s&txt_courseNumber=%s"
