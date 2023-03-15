package tv.sweet.testtask.extensions

import androidx.fragment.app.Fragment
import androidx.lifecycle.Lifecycle
import androidx.lifecycle.lifecycleScope
import androidx.lifecycle.repeatOnLifecycle
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.launch

fun <T, F : Flow<T>> Fragment.collectFlow(flow: F, action: (T) -> Unit) {
    scope().launch {
        repeatOnLifecycle(Lifecycle.State.STARTED) {
            flow.collect {
                it?.let(action)
            }
        }
    }
}

fun Fragment.scope() = viewLifecycleOwner.lifecycleScope