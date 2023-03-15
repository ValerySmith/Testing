package tv.sweet.testtask.core

import android.os.Bundle
import androidx.annotation.CallSuper
import androidx.annotation.LayoutRes
import androidx.appcompat.app.AppCompatActivity

abstract class BaseActivity(
    @LayoutRes private val layoutResourceId: Int
) : AppCompatActivity(), BaseView {
    @CallSuper
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        if (layoutResourceId != 0) {
            setContentView(layoutResourceId)
        }

        initializeViews()
        onObserveData()
    }

    abstract fun initializeViews()

    protected open fun onObserveData() {}

}