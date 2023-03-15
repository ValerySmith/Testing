package tv.sweet.testtask.data.service

import com.ua.mytrinity.tv_client.proto.MovieServiceOuterClass
import retrofit2.Response
import retrofit2.http.*
import tv.sweet.signup_service.SignupServiceOuterClass
import tv_service.TvServiceOuterClass

interface ApiService {

    @POST("SignupService/SetPhone")
    suspend fun setPhone(@Body request: SignupServiceOuterClass.SetPhoneRequest): Response<SignupServiceOuterClass.SetPhoneResponse>

    @POST("SignupService/SetCode")
    suspend fun setCode(@Body request: SignupServiceOuterClass.SetCodeRequest): Response<SignupServiceOuterClass.SetCodeResponse>

    @POST("TvService/GetChannels")
    suspend fun getChannels(@Body request: TvServiceOuterClass.GetChannelsRequest): Response<TvServiceOuterClass.GetChannelsResponse>

    @POST("TvService/OpenStream")
    suspend fun openStream(@Body request: TvServiceOuterClass.OpenStreamRequest): Response<TvServiceOuterClass.OpenStreamResponse>

    @POST("TvService/CloseStream")
    suspend fun closeStream(@Body request: TvServiceOuterClass.CloseStreamRequest): Response<TvServiceOuterClass.CloseStreamResponse>

    @POST("/MovieService/GetConfiguration")
    suspend fun getConfiguration(@Body request: MovieServiceOuterClass.GetConfigurationRequest): Response<MovieServiceOuterClass.GetConfigurationResponse>

    @POST("/MovieService/GetGenreMovies")
    suspend fun getGenreMovies(@Body request: MovieServiceOuterClass.GetGenreMoviesRequest): Response<MovieServiceOuterClass.GetGenreMoviesResponse>

    @POST("/MovieService/GetMovieInfo")
    suspend fun getMovieInfo(@Body request: MovieServiceOuterClass.GetMovieInfoRequest): Response<MovieServiceOuterClass.GetMovieInfoResponse>

    @POST("/MovieService/GetLink")
    suspend fun getLink(@Body request: MovieServiceOuterClass.GetLinkRequest): Response<MovieServiceOuterClass.GetLinkResponse>


}