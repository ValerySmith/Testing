package tv.sweet.testtask.presentation

import androidx.navigation.fragment.NavHostFragment
import dagger.hilt.android.AndroidEntryPoint
import tv.sweet.testtask.R
import tv.sweet.testtask.core.BaseActivity

@AndroidEntryPoint
class MainActivity : BaseActivity(R.layout.activity_main) {

    override fun initializeViews() {
        initNavigation()
    }

    private fun initNavigation() {
        supportFragmentManager.findFragmentById(R.id.fragment_container_main)
                as NavHostFragment
    }

}