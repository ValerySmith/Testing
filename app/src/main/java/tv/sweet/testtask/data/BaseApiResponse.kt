package tv.sweet.testtask.data

import retrofit2.Response
import tv.sweet.testtask.domain.errors.BaseError
import tv.sweet.testtask.domain.errors.NetworkError
import tv.sweet.testtask.domain.errors.UnknownError
import java.net.UnknownHostException

abstract class BaseApiResponse {

    suspend fun <T> safeApiCall(apiCall: suspend () -> Response<T>): NetworkResult<T> {
        try {
            val response = apiCall()
            return if (response.isSuccessful) {
                val body = response.body()
                NetworkResult.Success(body)
            } else {
                error(BaseError)
            }
        } catch (e: Exception) {
            return when (e) {
                is UnknownHostException -> {
                    error(NetworkError)
                }
                else -> {
                    error(UnknownError)
                }
            }
        }
    }

    private fun <T> error(error: ErrorType): NetworkResult<T> =
        NetworkResult.Error(error)
}