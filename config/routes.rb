Rails.application.routes.draw do
  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
  get 'get_problem_1_data', to: 'problems1#index'
  get 'get_problem_2_data', to: 'problems2#index'
end
