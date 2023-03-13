package tv.sweet.testtask.extensions

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import kotlinx.coroutines.*

fun ViewModel.launchWithContext(
    dispatchers: CoroutineDispatcher = Dispatchers.IO,
    action: suspend CoroutineScope.() -> Unit,
) = viewModelScope.launch {
    withContext(dispatchers) {
        action.invoke(this)
    }
}

