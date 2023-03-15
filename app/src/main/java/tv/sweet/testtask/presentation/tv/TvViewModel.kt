package tv.sweet.testtask.presentation.tv

import android.util.Log
import com.ua.mytrinity.tv_client.proto.MovieServiceOuterClass
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.flow.first
import tv.sweet.testtask.core.BaseViewModel
import tv.sweet.testtask.data.datastore.DataStore
import tv.sweet.testtask.data.onError
import tv.sweet.testtask.data.onSuccess
import tv.sweet.testtask.data.repo.MovieServiceRepository
import tv.sweet.testtask.data.repo.TvServiceRepository
import tv.sweet.testtask.extensions.launchWithContext
import tv_service.ChannelOuterClass
import tv_service.TvServiceOuterClass
import javax.inject.Inject

@HiltViewModel
class TvViewModel @Inject constructor(
    private val tvServiceRepository: TvServiceRepository,
    private val movieServiceRepository: MovieServiceRepository,
    private val dataStore: DataStore
) : BaseViewModel() {

    private val _channels = MutableStateFlow<List<ChannelOuterClass.Channel>?>(null)
    val channels get() = _channels.asStateFlow()

    fun getChannels() = launchWithContext {
        showLoading(true)
        val token = dataStore.authTokenFlow.first()
        if (token.isEmpty()) return@launchWithContext
        val request = TvServiceOuterClass.GetChannelsRequest.newBuilder()
            .setAuth(token)
            .build()
        tvServiceRepository.getChannels(request)
            .onSuccess {
                showLoading(false)
                _channels.value = it?.listList
            }.onError {
                showLoading(false)
            }
    }

    fun openStream(channelID: Int) = launchWithContext {
        val token = dataStore.authTokenFlow.first()
        val request = TvServiceOuterClass.OpenStreamRequest.newBuilder()
            .setAuth(token)
            .setChannelId(channelID)
            .build()
        tvServiceRepository.openStream(request)
            .onSuccess {
                Log.e("TEST", "stream - $it")
            }.onError {

            }
    }

    fun getConfiguration() = launchWithContext {
        val token = dataStore.authTokenFlow.first()
        val request =
            MovieServiceOuterClass.GetConfigurationRequest.newBuilder().setAuth(token).build()
        movieServiceRepository.getConfiguration(request)
            .onSuccess {

            }
            .onError {

            }
    }
}