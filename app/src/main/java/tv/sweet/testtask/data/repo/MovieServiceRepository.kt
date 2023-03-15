package tv.sweet.testtask.data.repo

import com.ua.mytrinity.tv_client.proto.MovieServiceOuterClass
import tv.sweet.testtask.data.BaseApiResponse
import tv.sweet.testtask.data.NetworkResult
import tv.sweet.testtask.data.service.ApiService
import tv.sweet.testtask.domain.repo.IMovieServiceRepository
import javax.inject.Inject

class MovieServiceRepository @Inject constructor(private val apiService: ApiService):
    BaseApiResponse(), IMovieServiceRepository {

    override suspend fun getConfiguration(request: MovieServiceOuterClass.GetConfigurationRequest): NetworkResult<MovieServiceOuterClass.GetConfigurationResponse> =
        safeApiCall { apiService.getConfiguration(request) }

    override suspend fun getGenreMovies(request: MovieServiceOuterClass.GetGenreMoviesRequest): NetworkResult<MovieServiceOuterClass.GetGenreMoviesResponse> =
        safeApiCall { apiService.getGenreMovies(request) }


    override suspend fun getMovieInfo(request: MovieServiceOuterClass.GetMovieInfoRequest): NetworkResult<MovieServiceOuterClass.GetMovieInfoResponse> =
        safeApiCall { apiService.getMovieInfo(request) }

    override suspend fun getLink(request: MovieServiceOuterClass.GetLinkRequest): NetworkResult<MovieServiceOuterClass.GetLinkResponse> =
        safeApiCall { apiService.getLink(request) }
}