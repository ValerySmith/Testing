package tv.sweet.testtask.presentation.login

import android.util.Log
import com.ua.mytrinity.tv_client.proto.Application
import com.ua.mytrinity.tv_client.proto.Device
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.channels.Channel
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.receiveAsFlow
import tv.sweet.signup_service.SignupServiceOuterClass
import tv.sweet.testtask.core.BaseViewModel
import tv.sweet.testtask.data.onError
import tv.sweet.testtask.data.onSuccess
import tv.sweet.testtask.data.repo.SignUpRepository
import tv.sweet.testtask.domain.repo.ISignUpRepository
import tv.sweet.testtask.extensions.launchWithContext
import javax.inject.Inject

@HiltViewModel
class LoginViewModel @Inject constructor(private val repo: SignUpRepository) : BaseViewModel() {

    private val _phone = MutableStateFlow("")

    private val _event = Channel<Event>()
    val event get() = _event.receiveAsFlow()

    sealed class Event {
        object NavigateToSendCode : Event()
        object NavigateToTvShow : Event()
    }

    fun sendPhone(phone: String) {
        var isValid = false
        if (phone.isNotEmpty()) {
            isValid = true
        }
        if (isValid) {

            val request =
                SignupServiceOuterClass.SetPhoneRequest.newBuilder()
                    .setPhone(phone)
                    .setDevice(
                        Device.DeviceInfo.newBuilder()
                            .setType(Device.DeviceInfo.DeviceType.DT_Android_Player)
                            .setApplication(
                                Application.ApplicationInfo.newBuilder()
                                    .setType(Application.ApplicationInfo.ApplicationType.AT_SWEET_TV_Player)
                                    .build()
                            )
                            .build()
                    ).build()
            launchWithContext {
                showLoading(true)
                repo.setPhone(request)
                    .onSuccess {
                        showLoading(false)
                        _phone.value = phone
                        _event.send(Event.NavigateToSendCode)
                        Log.e("TEST", "data - ${it?.status?.number}")
                    }
                    .onError {
                        showLoading(false)
                    }
            }
        }
    }

    fun sendCode(code: String) {
        if (code.isNotEmpty()) {
            val request = SignupServiceOuterClass
                .SetCodeRequest.newBuilder()
                .setPhone(_phone.value)
                .setConfirmationCode(code.toInt())
                .build()
            launchWithContext {
                repo.setCode(request)
                    .onSuccess {
                        _event.send(Event.NavigateToTvShow)
                    }
                    .onError {

                    }
            }
        }


    }
}