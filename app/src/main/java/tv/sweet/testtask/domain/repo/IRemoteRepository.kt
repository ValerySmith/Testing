package tv.sweet.testtask.domain.repo

import tv.sweet.signup_service.SignupServiceOuterClass
import tv.sweet.testtask.data.NetworkResult

interface IRemoteRepository {

    suspend fun setPhone(request: SignupServiceOuterClass.SetPhoneRequest): NetworkResult<SignupServiceOuterClass.SetPhoneResponse>
}