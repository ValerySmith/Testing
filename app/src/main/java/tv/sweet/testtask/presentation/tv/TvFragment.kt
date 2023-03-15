package tv.sweet.testtask.presentation.tv

import androidx.fragment.app.viewModels
import by.kirich1409.viewbindingdelegate.viewBinding
import com.google.android.exoplayer2.ExoPlayer
import com.google.android.material.tabs.TabLayoutMediator
import dagger.hilt.android.AndroidEntryPoint
import tv.sweet.testtask.R
import tv.sweet.testtask.core.BaseFragment
import tv.sweet.testtask.databinding.FragmentTvBinding

@AndroidEntryPoint
class TvFragment : BaseFragment(R.layout.fragment_tv) {

    private val binding: FragmentTvBinding by viewBinding()
    private val viewModel by viewModels<TvViewModel>()

    private var player: ExoPlayer? = null

    override fun initializeView() {
        initViewPager()
        initPlayer()
    }

    override fun onObserveData() {
        super.onObserveData()
    }

    private fun initViewPager() {
        binding.viewPager.apply {
            adapter = TvSlidePagerAdapter(this@TvFragment)
        }
        TabLayoutMediator(binding.tabs, binding.viewPager) { tab, pos ->
        }.attach()
    }

    private fun initPlayer() {
        player = ExoPlayer.Builder(requireContext())
            .build().also {
                binding.videoView.player = it
                it.prepare()
            }
    }

}