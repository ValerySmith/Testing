package tv.sweet.testtask.presentation.tv.channels

import android.os.Bundle
import android.view.View
import androidx.fragment.app.viewModels
import by.kirich1409.viewbindingdelegate.viewBinding
import dagger.hilt.android.AndroidEntryPoint
import tv.sweet.testtask.R
import tv.sweet.testtask.core.BaseFragment
import tv.sweet.testtask.databinding.FragmentChannelsBinding
import tv.sweet.testtask.databinding.FragmentTvBinding
import tv.sweet.testtask.extensions.collectFlow
import tv.sweet.testtask.presentation.tv.TvViewModel
import tv_service.ChannelOuterClass

@AndroidEntryPoint
class ChannelsFragment : BaseFragment(R.layout.fragment_channels) {

    private val viewModel by viewModels<TvViewModel>()
    private val binding: FragmentChannelsBinding by viewBinding()

    override fun initializeView() {

    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        viewModel.getChannels()
    }

    override fun onObserveData() {
        super.onObserveData()
        viewModel.withHandlers {
            collectFlow(this.channels) { data ->
                data?.let {
                    ChannelRvAdapter(
                        it
                    ) { channel ->
                        viewModel.openStream(channel.id)
                    }.also {
                        binding.recyclerViewChannels.adapter = it
                    }
                }
            }
        }
    }
}