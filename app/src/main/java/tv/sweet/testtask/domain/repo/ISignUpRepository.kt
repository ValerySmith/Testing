package tv.sweet.testtask.domain.repo

import tv.sweet.signup_service.SignupServiceOuterClass
import tv.sweet.testtask.data.NetworkResult

interface ISignUpRepository {

    suspend fun setPhone(request: SignupServiceOuterClass.SetPhoneRequest): NetworkResult<SignupServiceOuterClass.SetPhoneResponse>

    suspend fun setCode(request: SignupServiceOuterClass.SetCodeRequest): NetworkResult<SignupServiceOuterClass.SetCodeResponse>
}