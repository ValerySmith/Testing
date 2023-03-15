package tv.sweet.testtask.data.datastore

import android.content.Context
import android.util.Log
import androidx.datastore.core.DataStore
import androidx.datastore.preferences.core.Preferences
import androidx.datastore.preferences.core.edit
import androidx.datastore.preferences.core.emptyPreferences
import androidx.datastore.preferences.core.stringPreferencesKey
import androidx.datastore.preferences.preferencesDataStore
import dagger.hilt.android.qualifiers.ApplicationContext
import kotlinx.coroutines.flow.catch
import kotlinx.coroutines.flow.map
import java.io.IOException
import javax.inject.Inject

const val USER_PREF = "user_preferences"

val Context.dataStore: DataStore<Preferences> by preferencesDataStore(name = USER_PREF)

class DataStore @Inject constructor(@ApplicationContext context: Context) {

    private object PreferencesKeys {
        val PUSH_AUTH_TOKEN = stringPreferencesKey(AUTH_TOKEN)
        val PUSH_REFRESH_TOKEN = stringPreferencesKey(REFRESH_TOKEN)
    }

    private val dataStore = context.dataStore

    suspend fun setAuthToken(accessToken: String) {
        dataStore.edit {
            it[PreferencesKeys.PUSH_AUTH_TOKEN] = accessToken
        }
    }

    val authTokenFlow = dataStore.data
        .catch { exception ->
            if (exception is IOException) {
                emit(emptyPreferences())
            } else {
                throw exception
            }
        }.map {
            "Bearer ${it[PreferencesKeys.PUSH_AUTH_TOKEN]}"
        }

    private companion object {
        private const val AUTH_TOKEN = "accessToken"
        private const val REFRESH_TOKEN = "refreshToken"
    }
}