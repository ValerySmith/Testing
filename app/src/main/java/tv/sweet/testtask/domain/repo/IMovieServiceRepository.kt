package tv.sweet.testtask.domain.repo

import com.ua.mytrinity.tv_client.proto.MovieServiceOuterClass
import tv.sweet.testtask.data.NetworkResult

interface IMovieServiceRepository {

    suspend fun getConfiguration(request: MovieServiceOuterClass.GetConfigurationRequest): NetworkResult<MovieServiceOuterClass.GetConfigurationResponse>

    suspend fun getGenreMovies(request: MovieServiceOuterClass.GetGenreMoviesRequest): NetworkResult<MovieServiceOuterClass.GetGenreMoviesResponse>

    suspend fun getMovieInfo(request: MovieServiceOuterClass.GetMovieInfoRequest): NetworkResult<MovieServiceOuterClass.GetMovieInfoResponse>

    suspend fun getLink(request: MovieServiceOuterClass.GetLinkRequest): NetworkResult<MovieServiceOuterClass.GetLinkResponse>
}