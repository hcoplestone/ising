clear all;
close all;
clc;

system = 0;

n0 = 2000
n = 200
m = (200000 - n0)/n

AverageMagnetisations = [];
AverageMagnetisationsErrors = [];
AverageAbsoluteMagnetisations = [];
AverageAbsoluteMagnetisationsErrors = [];
AverageEnergies = [];
AverageEnergiesErrors = [];

Ts = [];

for file = dir('awesome/*.csv')'
    fname = ['./awesome/', file.name];
    data = csvread(fname, 1);

    Sweep = data(:,1);
    Temp = data(:,2);
    SubSystemID = data(:,3);
    Magnetisation = data(:,4);
    DimensionlessEnergy = data(:,5);
    
    Ts = [Ts Temp(1)]
        
    magnetisations = Magnetisation(n0:n:(n0+n*m));
    absolutemagnetisations = abs(magnetisations);
    energies = DimensionlessEnergy(n0:n:(n0+n*m));
    energies = energies ./ (40*40);
    
    AverageMagnetisations = [AverageMagnetisations mean(magnetisations)];
    AverageAbsoluteMagnetisations = [AverageAbsoluteMagnetisations mean(absolutemagnetisations)];
    AverageEnergies = [AverageEnergies mean(energies)];
    
    ErrorMagnetisation = std(magnetisations) / sqrt(length(magnetisations) - 1);
    ErrorAbsoluteMagnetisation = std(absolutemagnetisations) / sqrt(length(absolutemagnetisations) - 1);
    ErrorEnergy = std(energies) / sqrt(length(energies) - 1);

    AverageMagnetisationsErrors = [AverageMagnetisationsErrors ErrorMagnetisation];
    AverageAbsoluteMagnetisationsErrors = [AverageAbsoluteMagnetisationsErrors ErrorAbsoluteMagnetisation];
    AverageEnergiesErrors = [AverageEnergiesErrors ErrorEnergy];
end


figure;
% plot(Ts, AverageEnergies, '.-')
errorbar(Ts, AverageEnergies, AverageEnergiesErrors, '.')
xlabel('$T_0$', 'Interpreter', 'latex', 'FontSize', 16);
ylabel('$<\frac{E}{NJ}>$', 'Interpreter', 'latex', 'FontSize', 16);

figure;
% plot(Ts, AverageMagnetisations, '.-')
errorbar(Ts, AverageMagnetisations, AverageMagnetisationsErrors, '.')
xlabel('$T_0$', 'Interpreter', 'latex', 'FontSize', 16);
ylabel('$<\mathcal{M}>$', 'Interpreter', 'latex', 'FontSize', 16);

figure;
% plot(Ts, AverageMagnetisations, '.-')
errorbar(Ts, AverageAbsoluteMagnetisations, AverageAbsoluteMagnetisationsErrors, '.')
xlabel('$T_0$', 'Interpreter', 'latex', 'FontSize', 16);
ylabel('$<|\mathcal{M}|>$', 'Interpreter', 'latex', 'FontSize', 16);

% ylim([-1, 0.2])

% legend_handle = legend('-DynamicLegend');
% set(legend_handle,'Interpreter','latex')
% legend('show');

 
% figure;
% plot(Sweep, DimensionlessEnergy)
% ylabel('$E/J$', 'Interpreter', 'latex', 'FontSize', 16);