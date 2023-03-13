package tv.sweet.testtask.data

sealed class NetworkResult<T>(
    val data: T? = null,
    val errorType: ErrorType? = null
) {
    class Success<T>(data: T?) : NetworkResult<T>(data)
    class Error<T>(errorType: ErrorType, data: T? = null) : NetworkResult<T>(data, errorType)
}

inline fun <T : Any> NetworkResult<T>.onSuccess(action: (T?) -> Unit): NetworkResult<T> {
    if (this is NetworkResult.Success) action(data)
    return this
}

inline fun <T : Any> NetworkResult<T>.onError(action: (ErrorType?) -> Unit): NetworkResult<T> {
    if (this is NetworkResult.Error) action(errorType)
    return this
}