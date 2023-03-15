package tv.sweet.testtask.presentation.tv

import androidx.fragment.app.Fragment
import androidx.fragment.app.FragmentActivity
import androidx.viewpager2.adapter.FragmentStateAdapter
import tv.sweet.testtask.core.BaseFragment
import tv.sweet.testtask.presentation.tv.channels.ChannelsFragment
import tv.sweet.testtask.presentation.tv.movie.MovieFragment

class TvSlidePagerAdapter(fragment: Fragment): FragmentStateAdapter(fragment) {

    override fun getItemCount(): Int = 2

    override fun createFragment(position: Int): Fragment {
        when(position) {
            0 -> return ChannelsFragment()
            1 -> return MovieFragment()

        }
        return ChannelsFragment()
    }
}