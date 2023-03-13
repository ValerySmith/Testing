package tv.sweet.testtask.data.service

import retrofit2.Response
import retrofit2.http.Body
import retrofit2.http.GET
import retrofit2.http.POST
import tv.sweet.signup_service.SignupServiceOuterClass

interface ApiService {

    @POST("/SetPhone")
    suspend fun setPhone(@Body request: SignupServiceOuterClass.SetPhoneRequest): Response<SignupServiceOuterClass.SetPhoneResponse>
}