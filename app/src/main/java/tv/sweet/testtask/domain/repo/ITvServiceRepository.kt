package tv.sweet.testtask.domain.repo

import androidx.navigation.NavDeepLinkRequest
import tv.sweet.testtask.data.NetworkResult
import tv_service.TvServiceOuterClass

interface ITvServiceRepository {

    suspend fun getChannels(request: TvServiceOuterClass.GetChannelsRequest): NetworkResult<TvServiceOuterClass.GetChannelsResponse>

    suspend fun openStream(request: TvServiceOuterClass.OpenStreamRequest): NetworkResult<TvServiceOuterClass.OpenStreamResponse>

    suspend fun closeStream(request: TvServiceOuterClass.CloseStreamRequest): NetworkResult<TvServiceOuterClass.CloseStreamResponse>
}