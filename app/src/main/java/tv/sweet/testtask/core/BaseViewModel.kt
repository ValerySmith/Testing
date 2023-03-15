package tv.sweet.testtask.core

import androidx.lifecycle.ViewModel
import kotlinx.coroutines.flow.MutableSharedFlow
import kotlinx.coroutines.flow.asSharedFlow

abstract class BaseViewModel :ViewModel() {

    private val _events = MutableSharedFlow<Events>()
    val events get() = _events.asSharedFlow()

    sealed class Events {
        data class Error(val error: String) : Events()
        data class Loading(val state: Boolean) : Events()
    }

    suspend fun showLoading(state: Boolean) {
        _events.emit(Events.Loading(state))
    }
}