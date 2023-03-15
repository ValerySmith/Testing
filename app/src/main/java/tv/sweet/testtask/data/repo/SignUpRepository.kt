package tv.sweet.testtask.data.repo

import android.util.Log
import tv.sweet.signup_service.SignupServiceOuterClass
import tv.sweet.testtask.data.BaseApiResponse
import tv.sweet.testtask.data.NetworkResult
import tv.sweet.testtask.data.datastore.DataStore
import tv.sweet.testtask.data.service.ApiService
import tv.sweet.testtask.domain.repo.ISignUpRepository
import javax.inject.Inject

class SignUpRepository @Inject constructor(
    private val dataStore: DataStore,
    private val apiService: ApiService
) : BaseApiResponse(),
    ISignUpRepository {

    override suspend fun setPhone(request: SignupServiceOuterClass.SetPhoneRequest):
            NetworkResult<SignupServiceOuterClass.SetPhoneResponse> =
        safeApiCall {
            return@safeApiCall apiService.setPhone(request)
        }

    override suspend fun setCode(request: SignupServiceOuterClass.SetCodeRequest):
            NetworkResult<SignupServiceOuterClass.SetCodeResponse> =
        safeApiCall {
            val data = apiService.setCode(request).also { response ->
                Log.e("TEST", response.body().toString())
                response.body()?.authToken?.let {
                    if (it.isNotEmpty())
                        dataStore.setAuthToken(it)
                }
            }
            return@safeApiCall data
        }
}