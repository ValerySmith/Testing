package tv.sweet.testtask.presentation.login

import android.bluetooth.BluetoothClass
import android.util.Log
import androidx.lifecycle.ViewModel
import com.ua.mytrinity.tv_client.proto.Device
import com.ua.mytrinity.tv_client.proto.InternalMovieServiceGrpc
import dagger.hilt.android.lifecycle.HiltViewModel
import io.grpc.okhttp.OkHttpChannelBuilder
import okhttp3.OkHttpClient
import tv.sweet.signup_service.SignupServiceOuterClass
import tv.sweet.testtask.data.onError
import tv.sweet.testtask.data.onSuccess
import tv.sweet.testtask.data.repo.RemoteRepository
import tv.sweet.testtask.domain.repo.IRemoteRepository
import tv.sweet.testtask.extensions.launchWithContext
import javax.inject.Inject

@HiltViewModel
class LoginViewModel @Inject constructor(private val repo: RemoteRepository) : ViewModel() {

    fun sendPhone(phone: String) {
        var isValid = false
        if (phone.isNotEmpty()) {
            isValid = true
        }
        if (isValid) {
            val request =
                SignupServiceOuterClass.SetPhoneRequest.newBuilder().setPhone(phone)
                    .setDevice(
                        Device.DeviceInfo.newBuilder()
                            .setType(Device.DeviceInfo.DeviceType.DT_Android_Player).build()
                    ).build()
            launchWithContext {
                repo.setPhone(request)
                    .onSuccess {
                        Log.e("TEST", "data - ${it?.status?.number}")
                    }
                    .onError {

                    }
              //  InternalMovieServiceGrpc
            }
        }
    }
}