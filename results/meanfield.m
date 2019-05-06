% clear all;
% close all;
% clc;
% 
% data = csvread('meanfield_isingdata_CSV.csv', 1);
% 
% Temperature = data(:,1);
% Magnetisation = data(:,2);
% EnergyPerSite = data(:,3);
% SpecificHeat = data(:,4);
% MagneticSusceptibility = data(:,5);
% 
% T = 0:0.001:0.96
% T = [T 0.96:0.00000005:0.99]
% T = [T 0.99:0.000000005:1.0]
% X = 1 ./ T
% AnalyticMagnetisation = (1 - 1 ./ (sinh(X*log(1 + sqrt(2)))).^4).^(1/8);

figure;
hold on;
plot(Temperature, Magnetisation, '.-')
plot(T, AnalyticMagnetisation, '.')

xlabel('$T/T_c$', 'Interpreter', 'latex', 'FontSize', 17);
ylabel('$\mathcal{M}$', 'Interpreter', 'latex', 'FontSize', 17);
ax = gca;
ax.FontSize = 14; 

leg = legend('Mean field theory', 'Analytic')
leg.FontSize = 17
% 
% figure;
% plot(Temperature, EnergyPerSite, '.-')
% xlabel('$T/T_c$', 'Interpreter', 'latex', 'FontSize', 17);
% ylabel('$E/(NJ)$', 'Interpreter', 'latex', 'FontSize', 17);
% ax = gca;
% ax.FontSize = 14; 
% 
% figure;
% plot(Temperature, SpecificHeat, '.-')
% xlabel('$T/T_c$', 'Interpreter', 'latex', 'FontSize', 17);
% ylabel('$c/k_B$', 'Interpreter', 'latex', 'FontSize', 17);
% ax = gca;
% ax.FontSize = 14; 
% 
% figure;
% plot(Temperature, MagneticSusceptibility, '.-')
% xlabel('$T/T_c$', 'Interpreter', 'latex', 'FontSize', 17);
% ylabel('$\chi$', 'Interpreter', 'latex', 'FontSize', 17);
% ax = gca;
% ax.FontSize = 14; 

% hold off;
% xlabel('$\frac{T}{T_c}$', 'Interpreter', 'latex', 'FontSize', 16);
% % ylabel('$\mathcal{M}$', 'Interpreter', 'latex', 'FontSize', 16);
% % ylabel('$\frac{E}{NJ}$', 'Interpreter', 'latex', 'FontSize', 16);
% % ylabel('$\frac{c}{k_B}$', 'Interpreter', 'latex', 'FontSize', 16);
% ylabel('$\chi$', 'Interpreter', 'latex', 'FontSize', 16);
