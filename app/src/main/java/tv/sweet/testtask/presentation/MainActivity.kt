package tv.sweet.testtask.presentation

import androidx.core.view.isVisible
import androidx.navigation.fragment.NavHostFragment
import by.kirich1409.viewbindingdelegate.viewBinding
import dagger.hilt.android.AndroidEntryPoint
import tv.sweet.testtask.R
import tv.sweet.testtask.core.BaseActivity
import tv.sweet.testtask.databinding.ActivityMainBinding

@AndroidEntryPoint
class MainActivity : BaseActivity(R.layout.activity_main) {

    private val binding: ActivityMainBinding by viewBinding()

    override fun initializeViews() {
        initNavigation()
    }

    private fun initNavigation() {
        supportFragmentManager.findFragmentById(R.id.fragment_container_main)
                as NavHostFragment
    }

    override fun showLoader() {
        binding.progressBar.isVisible = true
    }

    override fun hideLoader() {
        binding.progressBar.isVisible = false
    }

}