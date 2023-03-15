package tv.sweet.testtask.data.repo

import tv.sweet.testtask.data.BaseApiResponse
import tv.sweet.testtask.data.NetworkResult
import tv.sweet.testtask.data.service.ApiService
import tv.sweet.testtask.domain.repo.ITvServiceRepository
import tv_service.TvServiceOuterClass
import javax.inject.Inject

class TvServiceRepository @Inject constructor(private val apiService: ApiService) :
    BaseApiResponse(), ITvServiceRepository {

    override suspend fun getChannels(request: TvServiceOuterClass.GetChannelsRequest):
            NetworkResult<TvServiceOuterClass.GetChannelsResponse> = safeApiCall {
        val data = apiService.getChannels(request)
        return@safeApiCall data
    }

    override suspend fun openStream(request: TvServiceOuterClass.OpenStreamRequest): NetworkResult<TvServiceOuterClass.OpenStreamResponse> {
        TODO("Not yet implemented")
    }

    override suspend fun closeStream(request: TvServiceOuterClass.CloseStreamRequest): NetworkResult<TvServiceOuterClass.CloseStreamResponse> {
        TODO("Not yet implemented")
    }
}