package tv.sweet.testtask.presentation.login

import android.os.Bundle
import android.view.View
import androidx.core.widget.doAfterTextChanged
import androidx.fragment.app.Fragment
import androidx.fragment.app.viewModels
import androidx.navigation.fragment.findNavController
import by.kirich1409.viewbindingdelegate.viewBinding
import dagger.hilt.android.AndroidEntryPoint
import tv.sweet.testtask.R
import tv.sweet.testtask.core.BaseFragment
import tv.sweet.testtask.databinding.FragmentLoginBinding
import tv.sweet.testtask.extensions.collectFlow

@AndroidEntryPoint
class LoginFragment : BaseFragment(R.layout.fragment_login) {

    private val viewModel: LoginViewModel by viewModels()
    private val binding by viewBinding<FragmentLoginBinding>()

    override fun initializeView() {
        binding.buttonLogin.setOnClickListener {
            viewModel.sendPhone(binding.textInputLayoutPhone.editText?.text.toString().trim())
        }
    }

    override fun onObserveData() {
        super.onObserveData()

        viewModel.withHandlers {
            collectFlow(event) { event ->
                when (event) {
                    LoginViewModel.Event.NavigateToSendCode ->
                        findNavController().navigate(R.id.action_loginFragment_to_sendCodeFragment)
                    else -> {}
                }
            }

        }
    }
}