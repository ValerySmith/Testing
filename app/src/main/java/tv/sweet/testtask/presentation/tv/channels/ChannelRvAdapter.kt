package tv.sweet.testtask.presentation.tv.channels

import android.view.LayoutInflater
import android.view.ViewGroup
import androidx.recyclerview.widget.RecyclerView
import androidx.recyclerview.widget.RecyclerView.ViewHolder
import com.bumptech.glide.Glide
import tv.sweet.testtask.databinding.ItemChannelBinding
import tv_service.ChannelOuterClass
import tv_service.TvServiceOuterClass

class ChannelRvAdapter(
    private val data: List<ChannelOuterClass.Channel>,
    private val listener: OnClickChannel
) :
    RecyclerView.Adapter<ChannelRvAdapter.ChannelViewHolder>() {

    fun interface OnClickChannel {
        fun onClick(data: ChannelOuterClass.Channel)
    }

    inner class ChannelViewHolder(private val binding: ItemChannelBinding) :
        ViewHolder(binding.root) {
        fun bind(item: ChannelOuterClass.Channel) {
            Glide.with(binding.imageViewIcon.context)
                .load(item.iconUrl)
                .into(binding.imageViewIcon)
            binding.textViewName.text = item.name
            binding.root.setOnClickListener {
                listener.onClick(item)
            }
        }
    }

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): ChannelViewHolder {
        val binding = ItemChannelBinding.inflate(LayoutInflater.from(parent.context), parent, false)
        return ChannelViewHolder(binding)
    }

    override fun getItemCount(): Int = data.size

    override fun onBindViewHolder(holder: ChannelViewHolder, position: Int) {
        holder.bind(data[position])
    }
}