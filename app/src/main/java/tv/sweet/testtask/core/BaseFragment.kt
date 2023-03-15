package tv.sweet.testtask.core

import android.os.Bundle
import android.view.View
import android.widget.Toast
import androidx.annotation.LayoutRes
import androidx.fragment.app.Fragment
import tv.sweet.testtask.extensions.collectFlow

abstract class BaseFragment(@LayoutRes private val resId: Int) : Fragment(resId) {

    private val activity by lazy {
        requireActivity() as BaseView
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        initializeView()
        onObserveData()
    }

    abstract fun initializeView()

    open fun onObserveData() = Unit

    open fun visibleLoading(state: Boolean) {
        if (state) activity.showLoader() else activity.hideLoader()
    }

    open fun displayError(error: String?) {
        Toast.makeText(requireContext(), error, Toast.LENGTH_SHORT).show()
    }

    inline fun <T : BaseViewModel> T.withHandlers(callback: T.() -> Unit) {
        callback(this)
        collectFlow(events) { event ->
            when (event) {
                is BaseViewModel.Events.Loading -> {
                    visibleLoading(event.state)
                }
                is BaseViewModel.Events.Error -> {
                    displayError(event.error)
                }
            }
        }
    }
}