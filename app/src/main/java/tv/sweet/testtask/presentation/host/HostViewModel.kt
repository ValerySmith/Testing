package tv.sweet.testtask.presentation.host

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext
import tv.sweet.signup_service.SignupServiceOuterClass
import javax.inject.Inject

@HiltViewModel
class HostViewModel @Inject constructor() : ViewModel() {

    private fun test() {
        viewModelScope.launch {
            withContext(Dispatchers.IO) {
                val setPhoneRequest =
                    SignupServiceOuterClass.SetPhoneRequest.newBuilder().setPhone("").build()

            }
        }
    }
}