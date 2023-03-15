package tv.sweet.testtask.presentation.login

import android.os.Bundle
import android.view.View
import androidx.fragment.app.Fragment
import androidx.fragment.app.activityViewModels
import androidx.navigation.fragment.findNavController
import by.kirich1409.viewbindingdelegate.viewBinding
import dagger.hilt.android.AndroidEntryPoint
import tv.sweet.testtask.R
import tv.sweet.testtask.core.BaseFragment
import tv.sweet.testtask.databinding.FragmentSendCodeBinding
import tv.sweet.testtask.extensions.collectFlow

@AndroidEntryPoint
class SendCodeFragment : BaseFragment(R.layout.fragment_send_code) {

    private val viewModel by activityViewModels<LoginViewModel>()
    private val binding: FragmentSendCodeBinding by viewBinding()

    override fun initializeView() {
        binding.buttonSendCode.setOnClickListener {
            viewModel.sendCode(binding.textInputLayoutCode.editText?.text.toString())
        }
    }

    override fun onObserveData() {
        super.onObserveData()
        viewModel.withHandlers {
            collectFlow(event) { event ->
                when (event) {
                    LoginViewModel.Event.NavigateToTvShow -> findNavController().navigate(R.id.action_sendCodeFragment_to_hostFragment)
                    else -> {}
                }
            }
        }

    }

}