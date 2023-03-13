package tv.sweet.testtask.data.repo

import retrofit2.Response
import tv.sweet.signup_service.SignupServiceOuterClass
import tv.sweet.testtask.data.BaseApiResponse
import tv.sweet.testtask.data.NetworkResult
import tv.sweet.testtask.data.service.ApiService
import tv.sweet.testtask.domain.repo.IRemoteRepository
import javax.inject.Inject

class RemoteRepository @Inject constructor(private val apiService: ApiService) : BaseApiResponse(),
    IRemoteRepository {

    override suspend fun setPhone(request: SignupServiceOuterClass.SetPhoneRequest): NetworkResult<SignupServiceOuterClass.SetPhoneResponse> =
        safeApiCall { apiService.setPhone(request) }
}